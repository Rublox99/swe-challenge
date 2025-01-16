# Email Search and Indexing Interface

## About the Project

The project is a monorepo that aims to provide a seamless interface for searching emails within an indexed database. Using the Enron email dataset, it offers a robust pipeline that indexes the dataset into a ZincSearch instance and provides a Vue-based frontend for easy querying. \
The repository is divided into three main sections:

1. **Indexer:** Golang program to index the Enron emails into ZincSearch.
2. **Backend:** Golang-powered backend to serve data and handle queries (Get All, Get With Filters & Get by ID).
3. **Frontend:** Vue.js-based interface for users to interact with the indexed data.


## Built With

![Vue.js](https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D)
![Vuetify](https://img.shields.io/badge/Vuetify-1867C0?style=for-the-badge&logo=vuetify&logoColor=AEDDFF)
![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)
![NodeJS](https://img.shields.io/badge/node.js-6DA55F?style=for-the-badge&logo=node.js&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![ZincSearch](https://img.shields.io/badge/zincsearch-%230377CC.svg?style=for-the-badge&logo=zincsearch&logoColor=green)
---

## Getting Started

### Prerequisites

1. **Download the Enron Corp Email Dataset:**

   - URL: [Enron Email Dataset](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)
   - Extract the `.tgz` file into your desired directory after downloading.

2. **Install Dependencies:**

   - Golang: [Install Go](https://golang.org/dl/)
   - Node.js (for frontend): [Install Node.js](https://nodejs.org/)

3. **Clone the Repository:**

   ```bash
   git clone https://github.com/Rublox99/swe-challenge.git
   cd swe-challenge
   ```

### Installation

1. **Set up ZincSearch Instance:**

    - Download the `.exe` or `binarie` from https://github.com/zincsearch/zincsearch/releases
    - Run a ZincSearch instance based on your desired OS (https://zincsearch-docs.zinc.dev/installation/)

    ```
    Runs predeterminate on http://localhost:4080
    ```
     
2. **ZincSearch Index Creation:**

    - Run the respective command in Bash/PowerShell or any desire CLI based on your OS  while ZincSearch engine is running:
    ```
    Linux: root/Mac_Linux_Index_Creation.txt
    ```
    ```
    Windows: root/Windows_Index_Creation.txt
    ```
     
3. **Indexer Setup:**

   - Navigate to the `indexer` directory:
     ```bash
     cd indexer
     ```
   - Install dependencies and run the indexer:
     ```bash
     go mod tidy
     go run main.go
     ```

#### REQUIRED: Change *const = mailsDir* (LINE:11 indexer/main.go) for your enron_mail directory ####
#### REQUIRED: Change *const (zincURL, username, password) (LINE:13 indexer/pkg/zinc/zinc.go) for their respective ZincSearch instance parameters ####

3. **Backend Setup:**

   - Navigate to the `backend` directory:
     ```bash
     cd backend
     ```
   - Install dependencies and start the backend:
     ```bash
     go mod tidy
     go run main.go
     ```

####  REQUIRED: Change *const (index, username, password)* (LINE:11 backend/controllers/search.controller.go) for their respective ZincSearch instance parameters

4. **Frontend Setup:**

   - Navigate to the `frontend` directory:
     ```bash
     cd frontend
     ```
   - Install dependencies and start the development server:
     ```bash
     npm install
     npm run dev
     ```

####  REQUIRED: Change *private static baseURL* (LINE:5 frontend/src/services/zincService.ts) to the backend respective URL
---

## Profiling the Indexer

Starts the profiling data inside the main function of main.go 
```
- go tool pprof -seconds <time> http://localhost:6060/debug/pprof/profile
- go tool pprof <binary> cpu.pprof
- sudo apt install graphviz (for Linux)
- (pprof) web opens an SVG chart <=> (pprof) svg > output.svg to save it as .svg
```

---

## Contact

For any question, feel free to reach out:

- **Ruben Almendares** - [camiloalmendarez47@gmail.com](mailto\:camiloalmendarez47@gmail.com)
- **Personal Portfolio** - [rubendev.online](https://rubendev.online)

