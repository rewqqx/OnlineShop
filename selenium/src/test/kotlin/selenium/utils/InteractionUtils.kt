package selenium.utils

import org.openqa.selenium.By
import org.openqa.selenium.Point
import org.openqa.selenium.WebDriver
import org.openqa.selenium.interactions.Actions

fun WebDriver.clickToPoint(point: Point) {
    Actions(this).moveByOffset(point.x, point.y)
        .doubleClick()
        .perform()

    Actions(this).moveByOffset(-point.x, -point.y)
        .perform()
}

fun WebDriver.dragFromPointToPoint(start: Point, end: Point) {
    Actions(this)
        .moveByOffset(start.x, start.y)
        .clickAndHold()
        .perform();

    Actions(this)
        .moveByOffset(end.x - start.x, end.y - start.y)
        .release()
        .perform()

    Actions(this)
        .moveByOffset(-end.x, -end.y)
        .perform()
}