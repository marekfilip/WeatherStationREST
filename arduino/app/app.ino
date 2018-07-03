#include <OneWire.h>
#include <DS18B20.h>

const byte ONEWIRE_PIN = 12;
const byte PHOTOCELL_PIN = 7;
byte _ONEWIRE_ADDRESS[8] = {0x28, 0xFF, 0x80, 0x34, 0x86, 0x16, 0x5, 0xF};

int photocellReading;

OneWire onewire(ONEWIRE_PIN);
DS18B20 sensors(&onewire);

void setup() {
  while(!Serial);
  Serial.begin(9600);

  sensors.begin(12);
  sensors.request(_ONEWIRE_ADDRESS);
}

void loop() {
  if (Serial.available() > 0) {
    do {
      Serial.read();
      delay(10);
    } while(Serial.available() > 0);

    Serial.print('^');
    Serial.print(getLightReading());
    Serial.print(';');
    Serial.print(getTemperature());
    Serial.println('$');
    delay(1000);
  }
}

int getLightReading() {
  return analogRead(PHOTOCELL_PIN);
}

float getTemperature() {
  if (sensors.available())
  {
    float temperature = sensors.readTemperature(_ONEWIRE_ADDRESS);
    sensors.request(_ONEWIRE_ADDRESS);
    return temperature;
  }

  return -9999;
}

// For init only
void getOnewireAddresses() {
  byte address[8];
  onewire.reset_search();
  while(onewire.search(address))
  {
    if (address[0] != 0x28)
      continue;

    if (OneWire::crc8(address, 7) != address[7])
    {
      Serial.println(F("Błędny adres, sprawdz polaczenia"));
      break;
    }

    for (byte i=0; i<8; i++)
    {
      Serial.print(F("0x"));
      Serial.print(address[i], HEX);

      if (i < 7)
        Serial.print(F(", "));
    }
    Serial.println();
  }

  while(1);
}
