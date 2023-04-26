import io.kotlintest.matchers.string.shouldContain
import org.openqa.selenium.WebDriver
import org.openqa.selenium.chrome.ChromeDriver
import org.openqa.selenium.chrome.ChromeOptions

fun main() {
    val options: ChromeOptions =
        ChromeOptions().addArguments(
            "--disable-dev-shm-usage",
            "--disable-gpu",
            "--no-sandbox",
            "--remote-allow-origins=*",
            "--headless"
        )
    val driver: WebDriver = ChromeDriver(options)

    driver.get("https://github.com/")
    driver.quit()
}