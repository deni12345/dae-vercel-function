const API_ENDPOINTS = [
  {
    id: "health",
    name: "Health Check",
    path: "/api/health",
    method: "GET",
  },
  {
    id: "sheet-listen",
    name: "Listen Sheet",
    path: "/api/sheet-listen",
    method: "GET",
  },
  {
    id: "collections",
    name: "Collections",
    path: "/api/collections",
    method: "GET",
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
    <div class="api-component">
       <div class="api-component-header">
          <button class="button" onclick="callApi('${endpoint.id}')">
             ${endpoint.method}
          </button>
        <div class="api-status">
          <h4 id="text-status-${endpoint.id}">Status: </h4>
          <span class="dot-status" id="dot-status-${endpoint.id}"></span>
        </div>
      </div>
        <div class="result-container">
          <div class="result-header" onclick="toggleResult('result-${endpoint.id}')">
            <span>Response</span>
            <span class="toggle-icon"></span>
          </div>
          <div id="result-${endpoint.id}" 
               class="result">
          </div>
          </div>
      </div>`;
    container.insertAdjacentHTML("beforeend", componentHtml);
  });
}

async function callApi(id) {
  const endpoint = API_ENDPOINTS.find((e) => e.id === id);
  const resultDiv = document.getElementById(`result-${id}`);

  try {
    resultDiv.classList.add("expanded");
    resultDiv.innerHTML = `<p>Loading...</p>`;
    const response = await axios.get(endpoint.path);
    resultDiv.innerHTML = `<pre>${JSON.stringify(
      response.data,
      null,
      2
    )}</pre>`;
    document.getElementById(
      `text-status-${id}`
    ).textContent = `Status: ${response.status}`;
    document.getElementById(`dot-status-${id}`).style.backgroundColor = "green";

    resultDiv.previousElementSibling.classList.add("expanded");
  } catch (error) {
    resultDiv.innerHTML = `<h3>Error:</h3><p>${error.message}</p>`;
    document.getElementById(`text-status-${id}`).textContent = `Status: Error`;
    document.getElementById(`dot-status-${id}`).style.backgroundColor = "red";
  }
}

document.addEventListener("DOMContentLoaded", generateApiComponents);
