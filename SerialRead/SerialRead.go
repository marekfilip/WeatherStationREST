package SerialRead

import (
	"bufio"
	"filip/tracer"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

type SerialRead struct {
	deviceFile *os.File
	deviceName string
	devicePath string
	mutex      *sync.Mutex
	tracer     tracer.Tracer
}

func Init() *SerialRead {
	var sr SerialRead = SerialRead{mutex: &sync.Mutex{}, tracer: tracer.Off()}
	sr.deviceName = sr.getUSBSerialDevice()
	sr.devicePath = "/dev/" + sr.deviceName

	var err error
	var sttyCommand *exec.Cmd = exec.Command("/bin/stty", []string{
		"-F",
		sr.devicePath,
		"cs8",
		"9600",
		"ignbrk",
		"-brkint",
		"-icrnl",
		"-imaxbel",
		"-opost",
		"-onlcr",
		"-isig",
		"-icanon",
		"-iexten",
		"-echo",
		"-echoe",
		"-echok",
		"-echoctl",
		"-echoke",
		"noflsh",
		"-ixon",
		"-crtscts",
	}...)

	err = sttyCommand.Run()
	sr.catchErrorAndExit(err)

	sr.deviceFile, err = os.OpenFile(sr.devicePath, syscall.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0666)
	sr.catchErrorAndExit(err)

	t := syscall.Termios{
		Iflag:  syscall.IGNPAR,
		Cflag:  syscall.CS8 | syscall.CREAD | syscall.CLOCAL | syscall.B9600,
		Cc:     [32]uint8{syscall.VMIN: 0, syscall.VTIME: uint8(20)}, //2.0s timeout
		Ispeed: syscall.B9600,
		Ospeed: syscall.B9600,
	}

	syscall.Syscall6(syscall.SYS_IOCTL, uintptr(sr.deviceFile.Fd()),
		uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&t)),
		0, 0, 0)

	return &sr
}

func (sr *SerialRead) getUSBSerialDevice() string {
	var err error
	var deviceName string
	var scannedText string
	var grepCmdReader io.ReadCloser
	var done chan bool = make(chan bool)
	var pattern string = "ch341-uart converter now attached to (ttyUSB[0-9]+)"
	var re *regexp.Regexp = regexp.MustCompile(pattern)
	var dmesgCmd *exec.Cmd = exec.Command("dmesg")
	var grepCmd *exec.Cmd = exec.Command("grep", "ttyUSB")

	var test bool

	grepCmd.Stdin, err = dmesgCmd.StdoutPipe()
	sr.catchErrorAndExit(err)

	grepCmdReader, err = grepCmd.StdoutPipe()
	sr.catchErrorAndExit(err)

	scanner := bufio.NewScanner(grepCmdReader)
	go func() {
		for scanner.Scan() {
			scannedText = scanner.Text()
			test, _ = regexp.MatchString(pattern, scannedText)
			if match, _ := regexp.MatchString(pattern, scannedText); match {
				deviceName = re.FindStringSubmatch(scanner.Text())[1]
			}
		}
		done <- true
	}()

	err = grepCmd.Start()
	sr.catchErrorAndExit(err)

	err = dmesgCmd.Run()
	sr.catchErrorAndExit(err)

	err = grepCmd.Wait()
	sr.catchErrorAndExit(err)

	<-done
	return deviceName
}

func (sr *SerialRead) catchErrorAndExitCustomMessage(msg string, err error) {
	if err != nil {
		log.Println(msg)
		log.Fatalln(err.Error())
	}
}

func (sr *SerialRead) catchErrorAndExit(err error) {
	sr.catchErrorAndExitCustomMessage("Could not start. Is device connected?", err)
}

func (sr *SerialRead) sendSignal() {
	var bytesSend int = 0
	var err error

	for bytesSend == 0 {
		bytesSend, err = sr.deviceFile.Write([]byte("a"))

		if err != nil {
			sr.catchErrorAndExitCustomMessage("Could not send signal.", err)
		}
	}
}

func (sr *SerialRead) ReadData() []byte {
	var byteCount int = 0
	var buffer []byte

	for byteCount < 8 {
		sr.lock()

		sr.sendSignal()
		<-time.After(time.Duration(200) * time.Millisecond)
		buffer = make([]byte, 16)
		byteCount, _ = sr.deviceFile.Read(buffer)

		sr.tracer.Trace("[", time.Now().Format("02-01-06 15:04:05"), "]", " Bytes read: ", byteCount, " Content: ", string(buffer))

		sr.unlock()
	}

	return buffer
}

func (sr *SerialRead) SetTracer(t tracer.Tracer) {
	sr.tracer = t
}

func (sr *SerialRead) lock() {
	sr.mutex.Lock()
}

func (sr *SerialRead) unlock() {
	sr.mutex.Unlock()
}
