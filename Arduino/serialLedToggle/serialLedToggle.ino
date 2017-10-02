//int led = 2;
int ledpin = LED_BUILTIN;
byte ledval = HIGH;
char Answer = 'a';

void setup() {
  Serial.begin(9600);
  pinMode(ledpin,OUTPUT);
}

void loop() {
  if (Serial.available()) {
    char inChar = (char)Serial.read();
    
    if (inChar == 'A') {
      ledval = HIGH; Answer = 'a';
    } else if (inChar == 'B'){
      ledval = LOW; Answer = 'b';
    } else if (inChar == 'Z'){
      ledval = LOW; Answer = 'z';
    }
    digitalWrite(ledpin,ledval);
    Serial.print(Answer);
    
  }
}

/*void serialEvent() {
  while (Serial.available()) {
    char inChar = (char)Serial.read();
    
    if (inChar == 'A') {
      digitalWrite(led,HIGH);
    } else if (inChar == 'B'){
      digitalWrite(led,LOW);
    }
  }
}*/
