package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 读取命令行参数 port，默认 6000
	port := flag.Int("port", 6000, "监听端口")
	flag.Parse()

	// 处理根路径，返回 HTML 页面
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, pageHTML)
	})

	// 启动服务器
	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("访问 http://localhost%s 查看密码生成器\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// HTML 页面（放在 Go 文件中，简化部署）
const pageHTML = `<!DOCTYPE html>
<html lang="zh">
<head>
<meta charset="UTF-8">
<title>强密码生成器</title>
<style>
body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: #f0f2f5;
  margin: 0;
}
.container {
  background: #fff;
  padding: 30px 40px;
  border-radius: 16px;
  box-shadow: 0 8px 20px rgba(0,0,0,0.1);
  text-align: center;
  width: 380px;
}
h1 { margin-bottom: 20px; font-size: 1.5rem; color: #333; }
input[type=number], input[type=text] { width: 60px; padding: 6px; margin-left: 10px; border-radius: 6px; border: 1px solid #ccc; }
label { display: block; margin: 10px 0 5px; font-size: 0.95rem; color: #555; }
button { background-color: #4caf50; color: white; padding: 10px 20px; margin-top: 15px; border: none; border-radius: 8px; cursor: pointer; font-size: 1rem; transition: 0.2s; }
button:hover { background-color: #45a049; }
#result-container { margin-top: 20px; display: flex; align-items: center; justify-content: space-between; background: #f5f5f5; padding: 10px; border-radius: 8px; word-break: break-all; }
#result { flex: 1; font-weight: bold; margin-right: 10px; }
#customCharsContainer { margin-top: 10px; display: none; }
#customChars { width: 100%; padding: 5px; border-radius: 6px; border: 1px solid #ccc; }
#strength { margin-top: 10px; font-weight: bold; }
.weak { color: red; }
.medium { color: orange; }
.strong { color: green; }
</style>
</head>
<body>
<div class="container">
<h1>强密码生成器</h1>

<label>
  密码长度：
  <input type="number" id="length" value="12" min="4" max="64">
</label>

<label>
  <input type="checkbox" id="symbols" checked>
  包含符号
</label>

<label>
  <input type="checkbox" id="enableCustom">
  启用自定义字符
</label>

<div id="customCharsContainer">
  <label>
    自定义字符：
    <input type="text" id="customChars" placeholder="如 ABC123@#">
  </label>
</div>

<button id="generate">生成密码</button>

<div id="result-container">
  <p id="result">点击生成密码</p>
  <button id="copy">复制</button>
</div>

<p id="strength">密码强度：未知</p>
</div>

<script>
const generateBtn = document.getElementById("generate");
const copyBtn = document.getElementById("copy");
const resultEl = document.getElementById("result");
const lengthInput = document.getElementById("length");
const symbolsCheckbox = document.getElementById("symbols");
const enableCustom = document.getElementById("enableCustom");
const customCharsContainer = document.getElementById("customCharsContainer");
const customCharsInput = document.getElementById("customChars");
const strengthEl = document.getElementById("strength");

enableCustom.addEventListener("change", function() {
  if (enableCustom.checked) {
    customCharsContainer.style.display = "block";
    symbolsCheckbox.disabled = true;
  } else {
    customCharsContainer.style.display = "none";
    symbolsCheckbox.disabled = false;
  }
});

function generatePassword(length, includeSymbols, customChars, useCustom) {
  const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  const symbols = "!@#$%^&*()_+[]{}|;:,.<>?";
  let chars = useCustom ? customChars : letters;
  if (!useCustom && includeSymbols) chars += symbols;
  if (!chars) return "";
  
  let password = "";
  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * chars.length);
    password += chars[randomIndex];
  }
  return password;
}

function getPasswordStrength(password) {
  let score = 0;
  if (!password) return {text: "未知", className: ""};
  if (password.length >= 8) score++;
  if (password.length >= 12) score++;
  if (/[a-z]/.test(password)) score++;
  if (/[A-Z]/.test(password)) score++;
  if (/[0-9]/.test(password)) score++;
  if (/[^a-zA-Z0-9]/.test(password)) score++;
  if (score <= 2) return {text: "弱", className: "weak"};
  if (score <= 4) return {text: "中", className: "medium"};
  return {text: "强", className: "strong"};
}

function updatePassword() {
  const length = parseInt(lengthInput.value);
  const useCustom = enableCustom.checked;
  const customChars = customCharsInput.value;
  const includeSymbols = symbolsCheckbox.checked;
  const password = generatePassword(length, includeSymbols, customChars, useCustom);
  resultEl.textContent = password;

  const strength = getPasswordStrength(password);
  strengthEl.textContent = "密码强度：" + strength.text;
  strengthEl.className = strength.className;
}

generateBtn.addEventListener("click", updatePassword);

copyBtn.addEventListener("click", function() {
  const password = resultEl.textContent;
  if (!password || password === "点击生成密码") return;
  navigator.clipboard.writeText(password).then(() => {
    alert("密码已复制到剪贴板！");
  });
});
</script>
</body>
</html>`
