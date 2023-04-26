package selenium.compare

import enums.EButtonTypes
import io.kotlintest.matchers.string.shouldContain
import io.kotlintest.shouldNotBe
import nu.pattern.OpenCV
import org.junit.jupiter.api.Test
import org.opencv.imgcodecs.Imgcodecs
import org.openqa.selenium.By
import org.openqa.selenium.WebDriver
import org.openqa.selenium.chrome.ChromeDriver
import org.openqa.selenium.chrome.ChromeOptions
import selenium.utils.*

class SeleniumCompareTest {
    private val options: ChromeOptions = getDriverOptions()
    private val driver: WebDriver = ChromeDriver(options)
    private val url: String = "http://157.230.77.165:8080/"

    init {
        OpenCV.loadShared()
        System.loadLibrary(org.opencv.core.Core.NATIVE_LIBRARY_NAME)
    }

    @Test
    fun compareJSONTest() {
        driver.get(url)

        val importJSONInput = driver.findButtonByType(EButtonTypes.UPLOAD_JSON_BUTTON)
        importJSONInput.shouldNotBe(null)
        importJSONInput!!.sendKeys(path + "/src/test/resources/jsons/rout_from_a_to_b.json")

        Imgcodecs.imwrite("./output/constructor-image.jpg", driver.getElementScreenshotMat("konvajs-content"))

        driver.tryClickElement(EButtonTypes.EXPORT_COLLAPSE)
        driver.tryClickElement(EButtonTypes.DOWNLOAD_HTML_BUTTON)

        val filePath = waitForFileDownload()
        driver.get("file:///" + filePath)
        val check = driver.findElement(By.id("konva_interactive_container"))
        check.shouldNotBe(null)

        Imgcodecs.imwrite("./output/generator-image.jpg", driver.getElementScreenshotMat("konvajs-content"))

        clearDownloadedFiles()

        driver.quit()
    }
}