# Mailgun Setup Guide (Developer)

This guide is for you (the developer) to set up Mailgun for KPK Accounting's contact form. Your client doesn't need to do anything — Mailgun is invisible to them.

## Overview

- **What it does:** Sends contact form emails to ken@kpkaccounting.com (their Outlook inbox)
- **Cost:** Free (5,000 emails/month included)
- **Your client:** Never needs to log in or manage anything
- **Time to complete:** ~15 minutes

---

## Step 1: Create Your Mailgun Account (2 minutes)

1. Go to [mailgun.com](https://mailgun.com)
2. Click "Sign Up Free"
3. Enter your email and create a password
4. Verify your email address
5. You now have a free account with 5,000 emails/month

---

## Step 2: Get Your API Credentials (2 minutes)

1. Log in to [app.mailgun.com](https://app.mailgun.com)
2. Click on "Send" in the left sidebar
3. Under "Sending > Domains", copy the domain that looks like: `mg.xxxxxxxxx.mailgun.org`
4. Click on that domain
5. When prompted to "Choose Method", select "API"
6. You'll see a message about setting up a "sending key" — click the link
7. Enter a description like "KPK Accounting Website" and save
8. Copy the API key that appears

**You now have:**
- Domain: `mg.xxxxxxxxx.mailgun.org`
- API Key: `key-xxxxxxxxxxxxxxx`

---

## Step 3: Add DNS Records to GoDaddy (5-10 minutes)

This tells GoDaddy's DNS to use Mailgun for sending emails from kpkaccounting.com.

1. Log in to [godaddy.com](https://godaddy.com)
2. Click your account → "Manage my products"
3. Find the kpkaccounting.com domain → Click "Manage DNS"
4. Back in Mailgun, go to your domain settings
5. Look for "DNS Records" or "DNS Verification"
6. Copy the DNS records Mailgun provides (usually 3-4 records)
7. In GoDaddy's DNS editor, add each record exactly as shown
8. Save and wait for DNS to propagate (usually instant, max 24 hours)

Once done, Mailgun will show a green checkmark indicating DNS is verified.

---

## Step 4: Update `.env` File (1 minute)

Replace the placeholders in your `.env` file with your actual credentials:

```env
MAILGUN_DOMAIN=mg.xxxxxxxxx.mailgun.org
MAILGUN_API_KEY=key-xxxxxxxxxxxxxxx
MAILGUN_FROM_EMAIL=ken@kpkaccounting.com
EMAIL_PROVIDER=mailgun
```

---

## Step 5: Test (1 minute)

1. Run your website: `go run ./cmd/web`
2. Fill out the contact form
3. Check ken@kpkaccounting.com's Outlook inbox
4. Email should arrive within seconds

---

## That's It!

Your client's contact form is now working. They receive emails in their normal Outlook inbox and don't need to know about Mailgun.

---

## Troubleshooting

**Email not arriving?**
- Check Mailgun dashboard for delivery logs
- Verify DNS records are correct in GoDaddy
- Check spam/junk folder

**Need to add more emails to receive?**
- Update the "To" email in handlers.go (currently hardcoded to ken@kpkaccounting.com)

**Exceeding 5,000 emails/month?**
- Upgrade to Mailgun's paid plan (~$35/month)

---

## API Key Security Note

⚠️ Keep your API key secret — treat it like a password.  
✓ Already in `.env` (which is in `.gitignore`)  
✓ Never commit `.env` to git  


