// 使用公钥进行加密
export async function rsaEncrypt(publicKeyPem: string, text: string) {
  const algorithm = { name: "RSA-OAEP", hash: "SHA-256" };

  const publicKey = await crypto.subtle.importKey(
    "spki",
    // 将 PEM 格式的公钥转换为 ArrayBuffer
    Uint8Array.from(
      atob(
        publicKeyPem
          .replace("-----BEGIN PUBLIC KEY-----", "")
          .replace("-----END PUBLIC KEY-----", "")
          .replace(/\s+/g, "")
      ),
      (c) => c.charCodeAt(0)
    ),
    algorithm,
    false,
    ["encrypt"]
  );

  const encrypted = await crypto.subtle.encrypt(
    algorithm,
    publicKey,
    new TextEncoder().encode(text)
  );

  return btoa(String.fromCharCode(...new Uint8Array(encrypted)));
}
