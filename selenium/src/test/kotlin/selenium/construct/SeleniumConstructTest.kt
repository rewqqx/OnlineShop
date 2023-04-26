package selenium.construct

import config.Config
import enums.EButtonTypes
import io.kotlintest.shouldBe
import io.kotlintest.shouldNotBe
import nu.pattern.OpenCV
import org.junit.jupiter.api.Test
import org.opencv.core.Mat
import org.opencv.core.MatOfByte
import org.opencv.imgcodecs.Imgcodecs
import org.openqa.selenium.*
import org.openqa.selenium.chrome.ChromeDriver
import org.openqa.selenium.chrome.ChromeOptions
import org.openqa.selenium.support.ui.Wait
import selenium.utils.*
import java.awt.Dimension
import java.nio.file.Paths


class SeleniumConstructTest {
    private val options: ChromeOptions = getDriverOptions()
    private val driver: WebDriver = ChromeDriver(options)
    private val url: String = "http://157.230.77.165:8080/"

    init {
        OpenCV.loadShared()
        System.loadLibrary(org.opencv.core.Core.NATIVE_LIBRARY_NAME)
    }

    private fun tryToMakeElementAndDragToOffset(type: EButtonTypes, offset: Point): List<Mat> {
        val beforeScreenshot = driver.getScreenshotMat()

        createElementByType(type)

        val afterScreenshot = driver.getScreenshotMat()
        val rect = afterScreenshot.findDifference(beforeScreenshot, null)

        driver.dragFromPointToPoint(rect.getCenterPoint(), rect.getCenterPointOffset(offset.x, offset.y))
        driver.clickToPoint(rect.getCenterPoint())

        val finalScreenshot = driver.getScreenshotMat()

        val screenshots = ArrayList<Mat>()
        screenshots.add(beforeScreenshot)
        screenshots.add(afterScreenshot)
        screenshots.add(finalScreenshot)

        return screenshots;
    }

    private fun createElementByType(type: EButtonTypes) {
        if (type != EButtonTypes.CREATE_IMAGE_BUTTON) {
            driver.tryClickElement(type)
        } else {
            downloadImageElement()
        }
    }

    private fun downloadImageElement() {
        val importJSONInput = driver.findButtonByType(EButtonTypes.CREATE_IMAGE_BUTTON)
        importJSONInput.shouldNotBe(null)
        importJSONInput!!.sendKeys(path + "/src/test/resources/images/hexagon.png")
    }


    @Test
    fun dragAndDropTest() {
        driver.get(url)

        driver.tryClickElement(EButtonTypes.ELEMENT_COLLAPSE)

        val screenshots = tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_RECT_BUTTON, Point(-100, -100))

        val rectBefore = screenshots[1].findDifference(screenshots[0], null)
        val rectAfter = screenshots[2].findDifference(screenshots[1], rectBefore)

        rectAfter.x.shouldBe(rectBefore.x - 100)
        rectAfter.y.shouldBe(rectBefore.y - 100)
        rectAfter.width.shouldBe(rectBefore.width)
        rectAfter.height.shouldBe(rectBefore.height)

        Imgcodecs.imwrite("./output/dragAndDropTest.jpg", driver.getScreenshotMat())

        driver.quit()
    }

    @Test
    fun allElementsTest() {
        driver.get(url)

        driver.tryClickElement(EButtonTypes.ELEMENT_COLLAPSE)

        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_RECT_BUTTON, Point(-360, -250))
        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_MARKER_BUTTON, Point(-240, -250))
        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_FRAME_BUTTON, Point(-120, -250))
        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_IMAGE_BUTTON, Point(0, -250))
        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_AREA_BUTTON, Point(120, -250))

        driver.tryClickElement(EButtonTypes.CONTROL_TAB)

        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_BUTTON_BUTTON, Point(-360, -75))
        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_INPUT_BUTTON, Point(-200, -75))
        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_SELECTOR_BUTTON, Point(80, -75))

        driver.tryClickElement(EButtonTypes.PUZZLE_TAB)

        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_CUT_BUTTON, Point(-360, 150))
        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_RECT_BUTTON, Point(-240, 150))
        tryToMakeElementAndDragToOffset(EButtonTypes.CREATE_IMAGE_BUTTON, Point(-120, 150))

        Imgcodecs.imwrite("./output/allElementsTest.jpg", driver.getScreenshotMat())

        driver.quit()
    }

}