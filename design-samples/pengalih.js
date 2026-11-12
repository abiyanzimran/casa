// Sample pages only (08 §19.12.11 item 2), never the product: the address carries the choice
// (?arah=limau|koral|biru, &tema=terang|gelap, &kontras=tinggi), nothing is stored on the device.
// Sets data-* on <html> before the page is drawn, points the manifest and icons at the chosen
// direction, and carries the choice into links to the other sample pages.
(() => {
  const q = new URLSearchParams(location.search);
  const root = document.documentElement;
  const pick = (name, allowed) => (allowed.includes(q.get(name)) ? q.get(name) : null);
  root.dataset.arah = pick('arah', ['limau', 'koral', 'biru']) ?? root.dataset.arah;
  const tema = pick('tema', ['terang', 'gelap']);
  if (tema) root.dataset.tema = tema;
  if (pick('kontras', ['tinggi'])) root.dataset.kontras = 'tinggi';

  const arah = root.dataset.arah;
  for (const link of document.querySelectorAll('link[rel="manifest"], link[rel~="icon"], link[rel="apple-touch-icon"]')) {
    link.href = link.getAttribute('href').replace('limau', arah);
  }

  const carried = new URLSearchParams({ arah });
  if (tema) carried.set('tema', tema);
  if (root.dataset.kontras) carried.set('kontras', 'tinggi');
  document.addEventListener('DOMContentLoaded', () => {
    for (const el of document.querySelectorAll('a[href$=".html"], iframe[src$=".html"]')) {
      const attr = el.tagName === 'A' ? 'href' : 'src';
      el.setAttribute(attr, `${el.getAttribute(attr)}?${carried}`);
    }
    // the index form shows the choice in force
    const now = { arah, tema: tema ?? '', kontras: root.dataset.kontras ?? '' };
    for (const input of document.querySelectorAll('form input[type="radio"]')) {
      input.checked = input.value === now[input.name];
    }
    // theme-color follows the page background when the theme is forced
    if (tema) {
      const background = getComputedStyle(root).backgroundColor;
      for (const meta of document.querySelectorAll('meta[name="theme-color"]')) {
        meta.removeAttribute('media');
        meta.content = background;
      }
    }
  });
})();
