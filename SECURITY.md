# Security Configuration Guide

## 🔒 **Required GitHub Secrets**

For secure deployment, add these secrets to your GitHub repository:

### **Repository Secrets** (`Settings > Secrets and variables > Actions`)

| Secret Name | Description | Example |
|-------------|-------------|---------|
| `GCP_PROJECT_ID` | Google Cloud Project ID | `ingka-find-racho8-dev` |
| `MICROSERVICE_URL` | Product Service URL | `https://product-service-xxx.run.app` |
| `GCP_SA_KEY` | Service Account JSON Key | `{"type": "service_account"...}` |
| `ALLOWED_ORIGIN` | Allowed CORS origin | `https://claude.ai` |

### **How to Set Up Secrets:**

<details>
<summary><strong>🔑 1. GCP_PROJECT_ID</strong></summary>

```bash
# Get your project ID
gcloud config get-value project

# Add to GitHub Secrets:
# Name: GCP_PROJECT_ID
# Value: your-project-id
```
</details>

<details>
<summary><strong>🌐 2. MICROSERVICE_URL</strong></summary>

```bash
# Get your microservice URL
gcloud run services list --filter="name:product-service"

# Add to GitHub Secrets:
# Name: MICROSERVICE_URL  
# Value: https://product-service-xxx.run.app
```
</details>

<details>
<summary><strong>🔐 3. GCP_SA_KEY</strong></summary>

```bash
# Create deployment service account
gcloud iam service-accounts create github-deployer \
  --description="Service account for GitHub Actions deployment"

# Grant necessary roles
gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member="serviceAccount:github-deployer@$PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/run.admin"

gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member="serviceAccount:github-deployer@$PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/storage.admin"

gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member="serviceAccount:github-deployer@$PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/artifactregistry.admin"

# Create and download key
gcloud iam service-accounts keys create github-deployer-key.json \
  --iam-account=github-deployer@$PROJECT_ID.iam.gserviceaccount.com

# Add entire JSON content to GitHub Secrets:
# Name: GCP_SA_KEY
# Value: {entire JSON content}
```
</details>

<details>
<summary><strong>🌍 4. ALLOWED_ORIGIN</strong></summary>

```bash
# For Claude AI (recommended)
ALLOWED_ORIGIN="https://claude.ai"

# For VS Code (if needed)
ALLOWED_ORIGIN="vscode-file://vscode-app"

# For local development
ALLOWED_ORIGIN="http://localhost:3000"

# Add to GitHub Secrets:
# Name: ALLOWED_ORIGIN
# Value: https://claude.ai
```
</details>

## 🛡️ **Security Features Implemented**

### **1. Environment Variables**
- ✅ All sensitive URLs moved to environment variables
- ✅ No hardcoded credentials in code
- ✅ GitHub Secrets for CI/CD

### **2. CORS Security**
- ✅ Restricted to specific origins (no wildcards)
- ✅ Disabled credentials for security
- ✅ Limited HTTP methods

### **3. Security Headers**
- ✅ `X-Content-Type-Options: nosniff`
- ✅ `X-Frame-Options: DENY`
- ✅ `X-XSS-Protection: 1; mode=block`
- ✅ `Strict-Transport-Security`
- ✅ `Content-Security-Policy`

### **4. Authentication**
- ✅ Google Cloud IAM with service accounts
- ✅ Bearer token authentication
- ✅ Cloud Run IAM policies

### **5. Logging Security**
- ✅ No sensitive URLs in logs
- ✅ Configuration status logging only
- ✅ Structured logging practices

## 🚨 **Security Checklist**

### **Before Deployment:**
- [ ] All secrets configured in GitHub
- [ ] Service account has minimal required permissions
- [ ] CORS origins properly restricted
- [ ] Environment variables not hardcoded
- [ ] `.gitignore` excludes all sensitive files

### **Regular Security Maintenance:**
- [ ] Rotate service account keys every 90 days
- [ ] Review IAM permissions quarterly
- [ ] Update dependencies regularly
- [ ] Monitor Cloud Run access logs
- [ ] Review CORS origins as needed

## 🔍 **Security Monitoring**

### **Google Cloud Monitoring:**
```bash
# Monitor authentication failures
gcloud logging read 'resource.type="cloud_run_revision" 
  AND severity>=ERROR 
  AND textPayload:"401"' --limit=10

# Monitor unusual access patterns
gcloud logging read 'resource.type="cloud_run_revision" 
  AND httpRequest.requestMethod="POST" 
  AND httpRequest.status!=200' --limit=10
```

### **Regular Security Audits:**
```bash
# Check IAM policies
gcloud run services get-iam-policy ravi-mcp-server --region=europe-west3

# Review service account permissions
gcloud projects get-iam-policy $PROJECT_ID \
  --filter="bindings.members:serviceAccount:*" \
  --format="table(bindings.role,bindings.members)"
```

## 📞 **Security Incident Response**

### **If Credentials are Compromised:**
1. **Immediately rotate service account keys**
2. **Review access logs for unauthorized usage**
3. **Update GitHub Secrets with new credentials**
4. **Redeploy service to clear any cached credentials**

### **Emergency Commands:**
```bash
# Disable service immediately
gcloud run services update ravi-mcp-server \
  --region=europe-west3 \
  --no-allow-unauthenticated

# Revoke all service account keys
gcloud iam service-accounts keys list \
  --iam-account=github-deployer@$PROJECT_ID.iam.gserviceaccount.com \
  --format="value(name)" | \
  xargs -I {} gcloud iam service-accounts keys delete {} \
  --iam-account=github-deployer@$PROJECT_ID.iam.gserviceaccount.com
```

## 📚 **Additional Security Resources**

- [Google Cloud Security Best Practices](https://cloud.google.com/security/best-practices)
- [Cloud Run Security Guide](https://cloud.google.com/run/docs/securing)
- [GitHub Actions Security](https://docs.github.com/en/actions/security-guides)
- [OWASP API Security Top 10](https://owasp.org/www-project-api-security/)
