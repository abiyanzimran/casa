// Sample pages only (08 §19.13.9 item 5), never the product: the address carries the theme and
// contrast choice (?tema=terang|gelap, &kontras=tinggi); nothing is stored on the device.
// Sets data-* on <html> before the page is drawn and carries the choice into links to the other
// sample pages and into the Catat background.
(() => {
  const q = new URLSearchParams(location.search);
  const root = document.documentElement;
  const tema = ['terang', 'gelap'].includes(q.get('tema')) ? q.get('tema') : null;
  if (tema) root.dataset.tema = tema;
  if (q.get('kontras') === 'tinggi') root.dataset.kontras = 'tinggi';

  const carried = new URLSearchParams();
  if (tema) carried.set('tema', tema);
  if (root.dataset.kontras) carried.set('kontras', 'tinggi');
  const suffix = carried.size ? `?${carried}` : '';

  document.addEventListener('DOMContentLoaded', () => {
    for (const a of document.querySelectorAll('a[href$=".html"]')) a.href = `${a.getAttribute('href')}${suffix}`;
    // the Catat background loads once, with the choice already in its address
    for (const frame of document.querySelectorAll('iframe[data-src]')) frame.src = `${frame.dataset.src}${suffix}`;
    // the index form shows the choice in force
    const now = { tema: tema ?? '', kontras: root.dataset.kontras ?? '' };
    for (const input of document.querySelectorAll('form input[type="radio"]')) {
      input.checked = input.value === now[input.name];
    }
    // a forced theme keeps only its own fixed theme-color
    if (tema) {
      const keep = tema === 'gelap' ? 'dark' : 'light';
      for (const meta of document.querySelectorAll('meta[name="theme-color"]')) {
        if (meta.media.includes(keep)) meta.removeAttribute('media');
        else meta.remove();
      }
    }
  });
})();
