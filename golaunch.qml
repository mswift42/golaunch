import QtQuick 2.1
import QtQuick.Controls 1.0

Item {
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
        // spacing: 8
        Text {
            text: "Enter Search Term"
            color: "#222222"
            font.pixelSize: 8
            x:4 
        }

        TextField {
            onAccepted: ctrl.search(text)
            id: textInput1
            width: 300
            focus: true
        }
        ListView {
            y: 14
            height: 260
            boundsBehavior: Flickable.DragAndOvershootBounds
            highlightRangeMode: ListView.ApplyRange
            spacing: 4
            width: parent.width
            model: ctrl.searchresult.len
            delegate: Rectangle {
                width:parent.width
                height:30 
                border.color: "#d0d0d0"
                color:"#e2e2e2"
                Column {
                    Row {
                        spacing:80
                        layoutDirection:Qt.RightToLeft
                            Text {
                                id: sresultname
                                text: ctrl.searchresult.name(index)
                                MouseArea {
                                    anchors.fill : parent
                                    onClicked: ctrl.select(sresult.text)
                                }
                        }
                            Text {
                                id: sresult
                                text: ctrl.searchresult.path(index)
                                anchors.rightMargin: 10
                                MouseArea {
                                    anchors.fill : parent
                                    onClicked: ctrl.select(sresult.text)
                                }
                        }
                    }
                }
            }

        }

    }
}
