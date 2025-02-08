# React + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh


# how to run
Here's a clear set of instructions for your team members to set up and start working on the project. You can add this to your `README.md` or share it with them directly.

---

# **Hackathon25-Fresh - Project Setup Guide 🚀**

## **Project Overview**
This is a **business dashboard** that allows users to:
- View how their posts are performing 📊  
- Add new posts ✍️  
- See analytics in one place 📈  

The project is built with:
- **Frontend:** React (using Vite)
- **Backend:** Go (Golang)
- **Version Control:** Git & GitHub

---

## **1️⃣ Clone the Repository**
Make sure you have Git installed. Then, open a terminal and run:

```bash
git clone git@github.com:amyrd/hackathon25-fresh.git
```
(or use HTTPS if you haven't set up SSH)
```bash
git clone https://github.com/amyrd/hackathon25-fresh.git
```

Then, navigate into the project folder:

```bash
cd hackathon25-fresh
```

---

## **2️⃣ Setup the Frontend (React)**
### **Install Node.js (if not installed)**
- Download and install [Node.js](https://nodejs.org/) (Recommended: LTS version).
- Verify installation:
  ```bash
  node -v
  npm -v
  ```

### **Install Dependencies**
```bash
npm install
```

### **Start the Frontend**
```bash
npm run dev
```
This should start the development server, and you can access the app at:
📌 [http://localhost:5173](http://localhost:5173)

---

## **3️⃣ Setup the Backend (Go)**
### **Install Go (if not installed)**
- Download and install [Go](https://go.dev/dl/).
- Verify installation:
  ```bash
  go version
  ```

### **Navigate to the Backend Folder**
```bash
cd backend
```

### **Initialize Go Module (Only Needed Once)**
If this is your first time setting up the backend:
```bash
go mod init backend
go mod tidy
```

### **Run the Backend Server**
```bash
go run main.go
```

By default, the backend should run on **http://localhost:8080**.

---

## **4️⃣ Setup Git for Collaboration**
Before making changes, create a new branch:

```bash
git checkout -b your-feature-branch
```

Make your changes, then commit and push:

```bash
git add .
git commit -m "Your commit message"
git push origin your-feature-branch
```

Then, open a **Pull Request (PR)** on GitHub for review.

---

## **5️⃣ Environment Variables (If Needed)**
If we use environment variables (e.g., for API keys), create a `.env` file:

```bash
touch .env
```
Then, add any necessary environment variables (we will update this as needed).

---

## **6️⃣ Useful Commands**
- **Check Git Status:** `git status`
- **Fetch Latest Code:** `git pull origin main`
- **Switch Branches:** `git checkout branch-name`
- **Merge Changes from Main:** `git merge main`
- **Stop Backend Server:** Press `Ctrl + C`

---

That’s it! 🎉 If you run into issues, message on our team chat.  
Let’s build something awesome! 🚀

---

Would you like me to add any details about **database setup**, **API routes**, or **project structure**?
