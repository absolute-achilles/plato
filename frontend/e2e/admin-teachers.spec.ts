import { expect, test } from "@playwright/test"

const uniqueEmail = () => `teacher-${Date.now()}@plato.edu`

test.describe("admin teachers", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("/sign-in")
    await page.fill('input[id="email"]',
      "admin@plato.local")
    await page.fill('input[id="password"]',
      "admin12345")
    await page.click('button[type="submit"]')
    await page.waitForURL("/")
  })

  test("admin can create a teacher", async ({ page }) => {
    const email = uniqueEmail()

    await page.goto("/admin/teachers")
    await page.click('button:has-text("Add Teacher")')

    await page.fill('input[id="username"]',
      "budi.santoso")
    await page.fill('input[id="email"]',
      email)
    await page.fill('input[id="password"]',
      "Password123")
    await page.fill('input[id="phone"]',
      "+62123456789")

    await page.click('button[role="combobox"][id="department"]')
    await page.click('[role="option"]:has-text("Mathematics")')

    await page.click('button[type="submit"]:has-text("Create Teacher")')

    await expect(page.locator(`text=${email}`)).toBeVisible()
    await expect(page.locator('text=Teacher created successfully')).toBeVisible()
  })
})
