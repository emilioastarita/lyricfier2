import QtQuick 2.2
import QtQuick.Window 2.2
import QtQuick.Controls 2.5
import QtQuick.Controls.Styles 1.3
import QtQuick.Layouts 1.12

MenuBar {
        Menu {
            title: "Settings"
            MenuItem {
                text: "+ Font size"
                onClicked: root.incFontFactor()
            }
            MenuItem {
                text: "- Font Size"
                onClicked: root.decFontFactor()
            }

            font {
                pixelSize: 11
            }
        }
        background: Rectangle {
              color:  "#0F1624"
        }
        delegate: MenuBarItem {
            id: menuBarItem
            contentItem: Text {
                text: menuBarItem.text
                font: menuBarItem.font
                opacity: enabled ? 1.0 : 0.3
                color: menuBarItem.highlighted ? "#ffffff" : "#f45b69"
                horizontalAlignment: Text.AlignLeft
                verticalAlignment: Text.AlignVCenter
                elide: Text.ElideRight
            }
            background: Rectangle {
                implicitWidth: 40
                implicitHeight: 40
                opacity: enabled ? 1 : 0.3
                color: menuBarItem.highlighted ? "#f45b69" : "transparent"
            }
        }
}