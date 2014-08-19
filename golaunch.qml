import QtQuick 2.1
import QtQuick.Controls 1.0

ApplicationWindow {
    width: 300
    height: 100
    Action {
        id: quitAction
        text: "&Quit"
        shortcut: "Ctrl+Q"
        onTriggered: ctrl.quit()
    }

    Column {
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.verticalCenter: parent.verticalCenter
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
            height: parent.height-y
            width: parent.width
            model: ctrl.searchresult.len
            delegate: Rectangle {
                height: 14
                Text {
                    text: ctrl.searchresult.text(index)
                    MouseArea {
                        anchors.fill : parent
                        onClicked: ctrl.select(index)
                    }
                }
            }

        }
    }
    }