function validarTelefone(telefone) {
  // Remove todos os caracteres não numéricos, exceto o "+"
  let numeroLimpo = telefone.replace(/\D/g, "");

  // Verifica se o número começa com "+" (indicando um código de país)

  // Adiciona o código do Brasil '55' se não estiver presente
  if (!numeroLimpo.startsWith("55")) {
    numeroLimpo = "55" + numeroLimpo;
  }

  // Remove o zero à esquerda do DDD, se existir (ex: 055 -> 55, 047 -> 47)
  if (numeroLimpo[2] === "0") {
    numeroLimpo = numeroLimpo.slice(0, 2) + numeroLimpo.slice(3);
  }

  // Adiciona o dígito 9 antes do número se não estiver presente (para números com 12 dígitos)
  if (numeroLimpo.length === 12) {
    numeroLimpo = numeroLimpo.slice(0, 4) + "9" + numeroLimpo.slice(4);
  }

  // Valida se o número final tem 13 dígitos (formato correto: 55 + DDD + 9 dígitos)
  return /^\d{13}$/.test(numeroLimpo)
    ? numeroLimpo
    : `Telefone Inválido: ${numeroLimpo}`;
}

// Testes
console.log(validarTelefone("62999243835"));
