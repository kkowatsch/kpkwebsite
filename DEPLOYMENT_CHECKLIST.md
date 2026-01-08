# Security Deployment Checklist

Before deploying to production, complete all items in this checklist to ensure maximum security for your client's website.

## CRITICAL (Must Complete Before Launch)

### HTTPS/TLS Configuration
- [ ] Obtain SSL/TLS certificate (use Let's Encrypt for free)
- [ ] Configure HTTPS on web server
- [ ] Redirect all HTTP traffic to HTTPS
- [ ] Set minimum TLS 1.2
- [ ] Configure certificate auto-renewal

### Environment Variables
- [ ] Create `.env` file with secure values:
  ```
  PORT=3000
  SMTP_USERNAME=your-email@gmail.com
  SMTP_PASSWORD=your-app-specific-password
  ```
- [ ] Verify `.env` is in `.gitignore` ✓ (already configured)
- [ ] Use app-specific password for Gmail (not your main password)
- [ ] Never commit `.env` to version control
- [ ] Store `.env` securely on server (restrict file permissions: `chmod 600`)

### Session Security
- [ ] Update `sessionManager.Cookie.Secure = true` (verify in main.go)
- [ ] Confirm SameSite=Lax is set
- [ ] Set appropriate Lifetime (currently 24 hours)
- [ ] Set appropriate IdleTimeout (currently 20 minutes)

### Security Headers
- [ ] Verify CSP policy allows all needed resources
- [ ] Test that security headers are present (curl -I https://yoursite.com)
- [ ] Verify HSTS is enabled for production domain
- [ ] Add domain to HSTS preload list: https://hstspreload.org

### CSRF Protection
- [ ] Test CSRF token validation on both contact forms
- [ ] Verify hidden CSRF field in forms
- [ ] Confirm tokens are unique per session
- [ ] Test that form without token fails

---

## HIGH Priority (Complete Before Production)

### Database & Data Storage
- [ ] Plan data retention policy for contact form submissions
- [ ] Decide where form submissions are stored
- [ ] Implement regular backups
- [ ] Define data deletion schedule
- [ ] Ensure backup security

### Email Security
- [ ] Test email delivery with real SMTP credentials
- [ ] Verify emails are encrypted in transit (TLS)
- [ ] Confirm client email address is receiving submissions
- [ ] Set up email forwarding if needed
- [ ] Test error scenarios (invalid SMTP, network issues)

### Input Validation
- [ ] Test form with various invalid inputs
- [ ] Verify validation messages are user-friendly
- [ ] Test special characters in all fields
- [ ] Verify HTML/JavaScript cannot be injected
- [ ] Test with extremely long inputs

### Error Logging
- [ ] Configure secure error logging (file or service)
- [ ] Verify logs don't contain sensitive information
- [ ] Set up log rotation (prevent disk from filling)
- [ ] Test that stack traces are logged but not exposed to users
- [ ] Review logs for any exposed credentials

---

## MEDIUM Priority (Complete Before Public Launch)

### Monitoring & Alerting
- [ ] Set up basic uptime monitoring
- [ ] Configure alerts for high error rates
- [ ] Monitor form submission rates (detect spam)
- [ ] Review logs daily for first week
- [ ] Set up automated log rotation

### Access Control
- [ ] Restrict server access to authorized personnel only
- [ ] Use SSH keys (not passwords) for server access
- [ ] Disable root login via SSH
- [ ] Configure firewall to allow only ports 80/443
- [ ] Whitelist IPs if client has limited user base

### Testing & Validation
- [ ] Run full end-to-end form submission test
- [ ] Test on multiple browsers (Chrome, Firefox, Safari, Edge)
- [ ] Test on mobile devices
- [ ] Verify all validation messages appear correctly
- [ ] Test confirmation page after submission
- [ ] Verify email is received from form submission

### Documentation
- [ ] Review SECURITY.md (provided)
- [ ] Document any custom configurations
- [ ] Create incident response plan
- [ ] Document backup/restore procedures
- [ ] Create user documentation if needed

---

## LOW Priority (Nice to Have)

### Performance & Optimization
- [ ] Enable caching headers for static assets
- [ ] Consider CDN for CSS/JS/fonts (Cloudflare, AWS CloudFront)
- [ ] Verify page load times are acceptable
- [ ] Test under load to ensure timeouts are appropriate

### Additional Security Layers
- [ ] Consider WAF (Cloudflare, AWS WAF)
- [ ] Implement rate limiting if spam becomes issue
- [ ] Add bot detection if needed (reCAPTCHA, hCaptcha)
- [ ] Enable DDoS protection
- [ ] Set up content security monitoring

### Compliance & Legal
- [ ] Review Privacy Policy (GDPR, CCPA compliance)
- [ ] Add Terms of Service if appropriate
- [ ] Document data retention policies
- [ ] Create data deletion procedures
- [ ] Document security measures for client assurance

---

## Post-Deployment Verification

### Week 1 (Daily)
- [ ] Check error logs daily
- [ ] Verify forms are working
- [ ] Monitor submission emails arriving
- [ ] Check server resources (CPU, memory, disk)
- [ ] Verify HTTPS is working on all pages

### Weekly (Every 7 days)
- [ ] Review error logs for patterns
- [ ] Check form submission volume
- [ ] Verify backups are running
- [ ] Test form submission end-to-end
- [ ] Monitor security headers (curl -I)

### Monthly (Every 30 days)
- [ ] Review access logs for suspicious activity
- [ ] Check for security advisories on dependencies
- [ ] Update dependencies if security patches available
- [ ] Review SMTP credentials rotation schedule
- [ ] Audit user access logs

### Quarterly (Every 90 days)
- [ ] Full security audit using OWASP ZAP
- [ ] Penetration testing (if budget allows)
- [ ] Review and update security policies
- [ ] Dependency security scan (`go list -u -m all`)
- [ ] Update this checklist based on lessons learned

---

## Deployment Commands

### Build for Production
```bash
cd /Users/kayleighschlupp/github/kpkwebsite
go build -o kpkwebsite ./cmd/web/
```

### Run Application
```bash
./kpkwebsite -addr=:3000
```

### Run with Environment Variables
```bash
export $(cat .env | xargs)
./kpkwebsite
```

### Using Systemd (Recommended for Linux)
```bash
# Create /etc/systemd/system/kpkwebsite.service
[Unit]
Description=KPK Website
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/var/www/kpkwebsite
ExecStart=/var/www/kpkwebsite/kpkwebsite
Restart=always

[Install]
WantedBy=multi-user.target

# Enable and start
sudo systemctl enable kpkwebsite
sudo systemctl start kpkwebsite
```

---

## Sign-Off

- **Prepared by**: [Your name]
- **Date**: [Current date]
- **Reviewed by**: [Client/Manager name]
- **Date Reviewed**: [Date]
- **Approved for Production**: [Approval date]

---

## Support & Questions

If you have any questions about security measures or need clarification on any items, refer to [SECURITY.md](./SECURITY.md) for detailed explanations.

For security issues discovered after deployment, follow the incident response plan documented in SECURITY.md.
