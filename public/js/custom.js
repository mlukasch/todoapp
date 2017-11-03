function setTime(timeout) {
  const element = document.querySelector('#timebox');
  setTimeout(() => {
    const t = new Date().toTimeString().replace(/.*(\d{2}:\d{2}:\d{2}).*/, '$1');
    element.innerHTML = t;
    setTime(1000);
  }, timeout);
}

window.onload = function () {
  setTime(0);
};

