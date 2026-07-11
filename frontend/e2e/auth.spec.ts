import { expect, test } from "@playwright/test"

test.describe("authentication", () => {
  test("admin can sign in and see the dashboard", async ({ page }) => {
    await page.goto("/sign-in")
    await page.fill('input[id="email"]',
      "admin@plato.local")
    await page.fill('input[id="password"]',
      "admin12345")
    await page.click('button[type="submit"]')

    await page.waitForURL("/")
    await expect(page.locator("text=Platform overview at a glance")).toBeVisible()
  })

  test("invalid credentials show an error", async ({ page }) => {
    await page.goto("/sign-in")
    await page.fill('input[id="email"]',
      "admin@plato.local")
    await page.fill('input[id="password"]',
      "wrong-password")
    await page.click('button[type="submit"]')

    await expect(page.locator("text=Invalid email or password")).toBeVisible()
  })
})
