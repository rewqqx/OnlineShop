package selenium.start

import enums.EButtonTypes
import io.kotlintest.matchers.string.shouldContain
import io.kotlintest.shouldNotBe
import org.junit.jupiter.api.Test
import org.openqa.selenium.By
import org.openqa.selenium.WebDriver
import org.openqa.selenium.chrome.ChromeDriver
import org.openqa.selenium.chrome.ChromeDriverService
import org.openqa.selenium.chrome.ChromeOptions
import selenium.utils.*
import java.io.File
import java.nio.file.Paths
import java.util.concurrent.TimeUnit


class SeleniumStartTest {

    private val url: String = "http://localhost:63344/OnlineShop/frontend/src/index.php?_ijt=4m0t3ult0grkiopiomhphkppmc&_ij_reload=RELOAD_ON_SAVE"

//    private val options: ChromeOptions = getDriverOptions()
//    private val driver: WebDriver = ChromeDriver(options)
    private val driverPath = "C:\\Users\\misha\\chrome\\chromedriver.exe"
    private val options: ChromeOptions = getDriverOptions()
    private val driver: WebDriver = ChromeDriver(ChromeDriverService.Builder().usingDriverExecutable(File(driverPath)).build(), options)

    init {
        driver.manage()?.timeouts()?.implicitlyWait(50, TimeUnit.MILLISECONDS)
        driver.manage()?.window()?.maximize()
    }


    @Test
    fun startShopTest(){
        driver.run{
            get(url)
            pageSource.shouldContain("Online shop")
        }
    }


    @Test
    fun checkItem(){
        driver.run{
            get(url)
            pageSource.shouldContain("Apple")
        }
    }

    @Test
    fun checkSigInTest() {
        driver.run {
            get(url)
            pageSource.shouldContain("Sign In")
            quit()
        }
    }

//    @Test
//    fun uploadJSONTest() {
//        driver.get(url)
//
//        val importJSONInput = driver.findButtonByType(EButtonTypes.UPLOAD_JSON_BUTTON)
//        importJSONInput.shouldNotBe(null)
//        importJSONInput!!.sendKeys(path + "/src/test/resources/jsons/rout_from_a_to_b.json")
//
//        driver.quit()
//    }
//
//    @Test
//    fun loadHTMLTest() {
//        driver.get(url)
//
//        val importJSONInput = driver.findButtonByType(EButtonTypes.UPLOAD_JSON_BUTTON)
//        importJSONInput.shouldNotBe(null)
//        importJSONInput!!.sendKeys(path + "/src/test/resources/jsons/rout_from_a_to_b.json")
//
//        driver.tryClickElement(EButtonTypes.EXPORT_COLLAPSE)
//        driver.tryClickElement(EButtonTypes.DOWNLOAD_HTML_BUTTON)
//
//        waitForFileDownload()
//        clearDownloadedFiles()
//
//        driver.quit()
//    }

    @Test
    fun openHTMLTest() {
        driver.get("file:///" + path + "/src/test/resources/html/my_interactive_picture.html");
        val check = driver.findElement(By.id("konva_interactive_container"))
        check.shouldNotBe(null)
        driver.quit()
    }


}