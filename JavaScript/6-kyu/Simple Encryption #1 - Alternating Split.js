function encrypt(text, n) {
    if (!text || n <= 0) return text;

    for (let i = 0; i < n; i++) {
        let odds = '';
        let evens = '';
        for (let j = 0; j < text.length; j++) {
            if (j % 2 === 0) {
                evens += text[j];
            } else {
                odds += text[j];
            }
        }
        text = odds + evens;
    }
    return text;
}

function decrypt(encryptedText, n) {
    if (!encryptedText || n <= 0) return encryptedText;

    const length = encryptedText.length;
    for (let i = 0; i < n; i++) {
        const mid = Math.floor(length / 2);
        let odds = encryptedText.slice(0, mid);
        let evens = encryptedText.slice(mid);
        let decryptedText = '';
        for (let j = 0; j < length; j++) {
            if (j % 2 === 0) {
                decryptedText += evens.charAt(j / 2);
            } else {
                decryptedText += odds.charAt(Math.floor(j / 2));
            }
        }
        encryptedText = decryptedText;
    }
    return encryptedText;
}
