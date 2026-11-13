// Sample pages only (08 §19.13.5 "Lembar"): the Catat sheet closes with Tutup, and on Ringkas also
// by dragging its handle or head down, and with Esc. Dragging is never the only way (WCAG 2.5.7). Closing plays
// the exit motion from tokens.css, then returns to Beranda with the theme and contrast choice.
(() => {
  const lembar = document.querySelector('.lembar');
  const seret = lembar.querySelector('.seret');
  const ringkas = matchMedia('(max-width: 599.98px)');
  let awal = null;
  let jari = null; // the pointer that started the drag; other pointers are ignored
  let geser = 0;

  const tutup = () => {
    lembar.style.removeProperty('transition');
    lembar.style.removeProperty('transform');
    lembar.close();
  };
  lembar.addEventListener('close', () => {
    setTimeout(() => location.assign(`beranda.html${location.search}`), 300);
  });
  lembar.querySelector('[data-tutup]').addEventListener('click', tutup);
  // the sample opens the sheet without showModal(), so Esc is handled here
  document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape' && lembar.open) tutup();
  });

  seret.addEventListener('pointerdown', (e) => {
    if (awal !== null || e.button !== 0 || !ringkas.matches || e.target.closest('button, select')) return;
    awal = e.clientY;
    jari = e.pointerId;
    geser = 0;
    seret.setPointerCapture(e.pointerId);
    lembar.style.transition = 'none';
  });
  seret.addEventListener('pointermove', (e) => {
    if (awal === null || e.pointerId !== jari) return;
    geser = Math.max(0, e.clientY - awal);
    lembar.style.transform = `translateY(${geser}px)`;
  });
  const lepas = (e) => {
    if (awal === null || e.pointerId !== jari) return;
    awal = null;
    if (geser > Math.min(120, lembar.offsetHeight / 4)) tutup();
    else {
      lembar.style.removeProperty('transition');
      lembar.style.removeProperty('transform');
    }
  };
  seret.addEventListener('pointerup', lepas);
  seret.addEventListener('pointercancel', lepas);
})();
