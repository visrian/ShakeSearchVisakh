// for paginating
let pageNumber = 0;
const Controller = {
  search: (ev) => {
    ev.preventDefault();
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    const response = fetch(`/search?q=${data.query}`).then((response) => {
      response.json().then((results) => {
        // set append to false so we overwrite the results instead of appending to existing results
        Controller.updateTable(false, results);
        // increment so that "Load More" returns the next page
        pageNumber += 1
      });
    });
  },

  updateTable: (append, results) => {
    const table = document.getElementById("table-body");
    const rows = [];
    for (let result of results) {
      rows.push(`<tr><td>${result}</td></tr>`);
    }
    // append or overwrite
    if (append == true) {
      table.innerHTML += rows;
    } else {
      table.innerHTML = rows;
    }
  },

  loadMore: (pageNumber) => {
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    // supply the page= param so that the /search request returns the next page
    const response = fetch(`/search?q=${data.query}&page=${pageNumber}`).then((response) => {
      response.json().then((results) => {
        // set append to True so we append to existing results instead of overwriting
        Controller.updateTable(true, results);
        // increment so that "Load More" returns the next page
        pageNumber += 1
      });
    });
  },
};


const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);

// look up the "Load More" button and add an event listener to it actually loads the next page of results
const loadMoreButton = document.getElementById("load-more");
loadMoreButton.addEventListener("click", () => {
    Controller.loadMore(pageNumber + 1);
});
