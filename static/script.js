const API_ENDPOINTS = [
  {
    id: "health",
    name: "Health Check",
    path: "/api/health",
  },
  {
    id: "sheet-subscribe",
    name: "Sheet Subscribe",
    path: "/api/sheet-subscribe",
  },
];

function toggleResult(id) {
  const resultDiv = document.getElementById(id);
  const header = resultDiv.previousElementSibling;
  resultDiv.classList.toggle("expanded");
  header.classList.toggle("expanded");
}

function generateApiComponents() {
  const container = document.getElementById("container");

  API_ENDPOINTS.forEach((endpoint) => {
    const componentHtml = `
        <button class="button" onclick="callApi('${endpoint.id}')">
          Get ${endpoint.name}
        </button>
        <div class="result-container">
          <div class="result-header" onclick="toggleResult('result-${endpoint.id}')">
            <span>${endpoint.name} Result</span>
            <span class="toggle-icon"></span>
          </div>
          <div id="result-${endpoint.id}" 
               class="result" 
               placeholder="Result will appear here...">
          </div>
        </div>`;
    container.insertAdjacentHTML("beforeend", componentHtml);
  });
}

async function callApi(id) {
  const endpoint = API_ENDPOINTS.find((e) => e.id === id);
  const resultDiv = document.getElementById(`result-${id}`);
  resultDiv.innerHTML = "";

  try {
    const response = await axios.get(endpoint.path);
    const dotId = `status-${id}`;
    resultDiv.innerHTML = `
        <h4>Status: ${response.status}<span id=${dotId} class="dot"></span></h4>
        <h3>Response:</h3>
        <pre>${JSON.stringify(response.data, null, 2)}</pre>`;
    document.getElementById(dotId).style.background = "green";

    resultDiv.classList.add("expanded");
    resultDiv.previousElementSibling.classList.add("expanded");
  } catch (error) {
    resultDiv.innerHTML = `<h3>Error:</h3><p>${error.message}</p>`;
  }
}

document.addEventListener("DOMContentLoaded", generateApiComponents);
