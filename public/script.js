async function shortenURL() {
    const input = document.getElementById('urlInput');
    const resultDiv = document.getElementById('result');
    const originalUrl = input.value;

    if (!originalUrl) {
        alert("Por favor escribe una URL");
        return;
    }

    try {
        const response = await fetch('/api/shorten', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ original_url: originalUrl })
        });

        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.error || 'Algo salió mal');
        }

        const host = window.location.origin; 
        const shortLink = `${host}/${data.id}`;

        resultDiv.style.display = 'block';
        resultDiv.innerHTML = `
            <p style="color:#94a3b8; margin:0 0 5px 0;">Tu link corto:</p>
            <a href="${shortLink}" target="_blank" style="font-size: 1.2rem; font-weight:bold;">${shortLink}</a>
        `;

    } catch (error) {
        resultDiv.style.display = 'block';
        resultDiv.innerHTML = `<p class="error">Error: ${error.message}</p>`;
    }
}