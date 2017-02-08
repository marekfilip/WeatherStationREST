package SerialRead

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
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
}

func Init() *SerialRead {
	var sr SerialRead = SerialRead{mutex: &sync.Mutex{}}
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
	sr.catchErrorCustomeMsgAndExit(err)

	sr.deviceFile, err = os.OpenFile(sr.devicePath, syscall.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0666)
	sr.catchErrorCustomeMsgAndExit(err)

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

	grepCmd.Stdin, err = dmesgCmd.StdoutPipe()
	sr.catchErrorCustomeMsgAndExit(err)

	grepCmdReader, err = grepCmd.StdoutPipe()
	sr.catchErrorCustomeMsgAndExit(err)

	scanner := bufio.NewScanner(grepCmdReader)
	go func() {
		for scanner.Scan() {
			scannedText = scanner.Text()
			if match, _ := regexp.MatchString(pattern, scannedText); match {
				deviceName = re.FindStringSubmatch(scanner.Text())[1]
			}
		}
		done <- true
	}()

	err = grepCmd.Start()
	sr.catchErrorCustomeMsgAndExit(err)

	err = dmesgCmd.Run()
	sr.catchErrorCustomeMsgAndExit(err)

	err = grepCmd.Wait()
	sr.catchErrorCustomeMsgAndExit(err)

	<-done
	return deviceName
}

func (sr *SerialRead) catchErrorAndExit(msg string, err error) {
	if err != nil {
		log.Println(msg)
		log.Println(err.Error())
		os.Exit(1)
	}
}

func (sr *SerialRead) catchErrorCustomeMsgAndExit(err error) {
	sr.catchErrorAndExit("Could not start. Is device connected?", err)
}

func (sr *SerialRead) GetData() map[string]string {
	var data string = string(sr.readData())
	var splitedData = strings.Split(data, ";")
	var output map[string]string = map[string]string{}

	for _, value := range splitedData {
		tmp := strings.Split(strings.TrimSpace(value), ":")

		if len(tmp) == 2 {
			output[tmp[0]] = tmp[1]
		}
	}

	return output
}

func (sr *SerialRead) sendSignal() {
	var bytesSend int = 0
	var err error

	for bytesSend == 0 {
		sr.Lock()
		bytesSend, err = sr.deviceFile.Write([]byte("a"))

		if err != nil {
			sr.catchErrorAndExit("Could not send signal.", err)
		}
		sr.Unlock()
	}
}

func (sr *SerialRead) readData() []byte {
	var byteCount int = 0
	var buffer []byte

	for byteCount < 8 {
		sr.sendSignal()
		sr.Lock()
		buffer = make([]byte, 16)
		byteCount, _ = sr.deviceFile.Read(buffer)
		sr.Unlock()
	}

	return buffer
}

func (sr *SerialRead) Lock() {
	sr.mutex.Lock()
}

func (sr *SerialRead) Unlock() {
	<-time.After(time.Duration(60) * time.Millisecond)
	sr.mutex.Unlock()
}
