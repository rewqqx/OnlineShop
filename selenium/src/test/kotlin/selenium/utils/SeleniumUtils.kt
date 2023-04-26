package selenium.utils

import enums.EButtonTypes
import io.kotlintest.shouldNotBe
import org.opencv.core.Mat
import org.opencv.core.MatOfByte
import org.opencv.imgcodecs.Imgcodecs
import org.openqa.selenium.*
import org.openqa.selenium.chrome.ChromeOptions
import java.awt.Dimension

fun getDriverOptions(): ChromeOptions {
    val windowSize = Dimension(1920, 1080)
    return ChromeOptions()
        .setExperimentalOption("prefs", mapOf("download.default_directory" to downloadDir))
        .addArguments(
            "--disable-dev-shm-usage",
            "--disable-gpu",
            "--no-sandbox",
            "--remote-allow-origins=*",
            "--headless=new",
            "window-size=${windowSize.width},${windowSize.height}"
        )
}

fun WebDriver.findButtonByType(type: EButtonTypes): WebElement? {
    return this.findElement(By.id(type.value))
}


fun WebDriver.getScreenshotMat(): Mat {
    val bytes = (this as TakesScreenshot).getScreenshotAs(OutputType.BYTES)
    return Imgcodecs.imdecode(MatOfByte(*bytes), Imgcodecs.IMREAD_UNCHANGED)
}


fun WebDriver.getElementScreenshotMat(className: String): Mat {
    val element = this.findElement(By.className(className))
    element.shouldNotBe(null)
    val bytes = (element as TakesScreenshot).getScreenshotAs(OutputType.BYTES)
    return Imgcodecs.imdecode(MatOfByte(*bytes), Imgcodecs.IMREAD_UNCHANGED)
}

fun WebDriver.tryClickElement(type: EButtonTypes) {
    val collapse = this.findButtonByType(type)
    collapse.shouldNotBe(null)
    collapse!!.click();
    Thread.sleep(1_000);
}