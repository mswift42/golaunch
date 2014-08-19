import QtQuick 2.1
import QtQuick.Controls 1.0

ApplicationWindow {
    width: 800
    height: 300
    Action {
        id: quitAction
        text: "&Quit"
        shortcut: "Ctrl+Q"
        onTriggered: ctrl.quit()
    }

    Column {
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.verticalCenter: parent.verticalCenter
        anchors.fill: parent
        spacing: 8
        Text {
            text: "Enter Search Term"
            color: "#222222"
            font.pixelSize: 8
            x:4 
        }

        TextField {
            onAccepted: ctrl.search(text)
            id: textInput1
            focus: true
        }
        ListView {
            y: 14
            height: 200
            spacing: 4
            width: parent.width
            model: ctrl.searchresult.len
            delegate: Rectangle {
                width:parent.width
                height:30 
                border.color: "black"
                color:"#e2e2e2"
                Text {
                    id: sresult
                    text: ctrl.searchresult.text(index)
                    MouseArea {
                        anchors.fill : parent
                        onClicked: ctrl.select(sresult.text)
                    }
                }
            }

        }
    }
    }
