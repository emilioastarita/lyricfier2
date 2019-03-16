import QtQuick 2.2
import QtQuick.Window 2.2
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.12

Item {
    Rectangle {
        color : "#0F1624"
        implicitWidth: parent.width
        implicitHeight: parent.height
        ColumnLayout {
            spacing: 5
            Text {
                    Layout.margins: 5
                    id: title
                    font { pixelSize: 18 }
                    text: song.title
                    color: "#f45b69"
            }
            Text {
                    Layout.leftMargin: 5
                    id: artist
                    font { pixelSize: 15 }
                    text: song.artist
                    color: "#028090"
            }
            TextEdit {
                id: lyric
                Layout.margins: 5
                color: "#e4fde1"
                text: song.lyric
                readOnly: true
                wrapMode: Text.WordWrap
                selectByMouse: true
                font { pixelSize: 13 }
            }
        }
    }
}
