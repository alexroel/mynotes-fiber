const dialog = document.getElementById("myModal");
const openButton = document.getElementById("openModal");
const closeButton = document.getElementById("closeModal");

openButton.addEventListener("click", () => {
  dialog.showModal();
});

closeButton.addEventListener("click", () => {
  dialog.close();
});

// Click fuera del modal para cerrarlo
dialog.addEventListener("click", (event) => {
  if (event.target === dialog) {
    dialog.close();
  }
});
