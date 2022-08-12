import processing.serial.*;

PImage front;
Serial port;

void setup() {
    // ここにポートを設定する
    String portStr = "";

    size(800, 600, P3D);
    background(#000000);

    if (portStr == "") {
        text("Please check console", 10, 30);
        front = loadImage("msg.png");
        image(front, 0, 0);
        String list[] = Serial.list();
        for (int i = 0; i < list.length; i++) {
            text(String.format("%d : %s", i, list[i]), 10, 120 + 12 * i);
            println(String.format("%d : %s", i, list[i]));
        }
    } else {
        port = new Serial(this, portStr, 115200);
        front = loadImage("tinygo.png");
        textureMode(IMAGE);
    }
}

float rotationX = 0;
float rotationY = 0;

void draw() {
    if (port != null) {
        background(0);
        translate(width / 2, height / 2);
        scale(0.4);

        if (port != null && port.available() > 0) {
            String str = port.readStringUntil('\n');
            //print(str);

            if (str != null && str.length() > 5) {
                String spl[] = split(str, " ");
                if (spl.length >= 2) {
                    rotationX = float(spl[0]);
                    rotationY = -float(spl[1]);
                }
            }
            println(rotationX, rotationY);
        }
        pushMatrix();
        rotateX(rotationX);
        rotateZ(rotationY);
        drawWioTerminal();
        popMatrix();
    }
}

void drawWioTerminal() {
    beginShape();
    texture(front);
    vertex(-500, -112, -500, 0, 700);    //V1
    vertex( 500, -112, -500, 700, 700);  //V2
    vertex( 500, -112, 500, 700, 0);     //V3
    vertex(-500, -112, 500, 0, 0);       //V4
    endShape();
}
