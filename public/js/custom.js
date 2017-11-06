function setTime(timeout) {
  const element = document.querySelector('#timebox');
  setTimeout(() => {
    const t = new Date().toTimeString().replace(/.*(\d{2}:\d{2}:\d{2}).*/, '$1');
    element.innerHTML = t;
    setTime(1000);
  }, timeout);
}

function sendRegisterForm() {
  const form = document.getElementById('registerForm');
  const formData = new FormData();
  for (const element of form.elements) {
    formData.append(element.name, element.value);
  }
  console.log('fetch http://localhost/new/user');
  fetch('/user/new', {
    method: 'POST',
    body: formData,
  }).then(response => response.json())
    .then(data => console.log(data));
}

window.onload = function () {
  setTime(0);
};

