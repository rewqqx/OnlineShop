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
import kotlin.test.DefaultAsserter.assertEquals
import kotlin.test.DefaultAsserter.assertTrue
import kotlin.test.assertEquals
import kotlin.test.assertTrue


class SeleniumStartTest {

    private val url: String = "http://localhost:9050/frontend/src/"

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

    @Test
    fun checkPhonesFilter() {
        driver.run {
            get(url)
            pageSource.shouldContain("Phones")
            quit()
        }
    }

    @Test
    fun checkHouseFilter() {
        driver.run {
            get(url)
            pageSource.shouldContain("House")
            quit()
        }
    }

    @Test
    fun checkFruitsFilter() {
        driver.run {
            get(url)
            pageSource.shouldContain("Fruits")
            quit()
        }
    }

    @Test
    fun checkEducationFilter() {
        driver.run {
            get(url)
            pageSource.shouldContain("Education")
            quit()
        }
    }

    @Test
    fun checkFoodFilter() {
        driver.run {
            get(url)
            pageSource.shouldContain("Food")
            quit()
        }
    }

    @Test
    fun checkClothesFilter() {
        driver.run {
            get(url)
            pageSource.shouldContain("Clothes")
            quit()
        }
    }

    @Test
    fun checkBuildingsFilter() {
        driver.run {
            get(url)
            pageSource.shouldContain("Buildings")
            quit()
        }
    }

    @Test
    fun checkElectronicsFilter() {
        driver.run {
            get(url)
            pageSource.shouldContain("Electronics")
            quit()
        }
    }

    @Test
    fun checkCarsFilter() {
        driver.run {
            get(url)
            pageSource.shouldContain("Cars")
            quit()
        }
    }

    @Test
    fun testDeepCars() {
        driver.run {
            get(url)
            val carsLink = findElement(By.linkText("Cars"))
            carsLink.click()
            assertEquals("http://localhost:9050/frontend/src/?tag_id=3", currentUrl)

            val productElements = findElements(By.className("item-card"))
            assertTrue(productElements.isNotEmpty())

            for (productElement in productElements) {
                val nameElement = productElement.findElement(By.className("item-desc"))
                assertTrue(nameElement.isDisplayed())

                val priceElement = productElement.findElement(By.className("item-image"))
                assertTrue(priceElement.isDisplayed())
            }

            quit()
        }
    }

    @Test
    fun testDeepFruits() {
        driver.run {
            get(url)
            val fruitLink = findElement(By.linkText("Fruits"))
            fruitLink.click()
            assertEquals("http://localhost:9050/frontend/src/?tag_id=8", currentUrl)

            val productElements = findElements(By.className("item-card"))
            assertTrue(productElements.isNotEmpty())

            for (productElement in productElements) {
                val nameElement = productElement.findElement(By.className("item-desc"))
                assertTrue(nameElement.isDisplayed())

                val priceElement = productElement.findElement(By.className("item-image"))
                assertTrue(priceElement.isDisplayed())
            }

            val itemLink = findElement(By.linkText("Apple - 1 \$"))
            itemLink.click()

            pageSource.shouldContain("Apple")


            quit()
        }
    }


    @Test
    fun openHTMLTest() {
        driver.get("file:///" + path + "/src/test/resources/html/my_interactive_picture.html");
        val check = driver.findElement(By.id("konva_interactive_container"))
        check.shouldNotBe(null)
        driver.quit()
    }

    //    @Test
//    fun testLogInCheck() {
//        driver.run {
//            get(url)
//            val loginLink = findElement(By.linkText("Sign In"))
//            loginLink.click()
//
//            val pageSource = pageSource
//            assertTrue(pageSource.contains("Password"), "Page should contain 'Password'")
//
//            quit()
//        }
//    }


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

}