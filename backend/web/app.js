const rows = document.querySelector('#rows');
const message = document.querySelector('#message');
async function load() {
  const response = await fetch('/api/v1/inspections');
  const data = await response.json();
  rows.innerHTML = data.items.map(item => `<tr><td>${item.id}</td><td>${item.site} / ${item.door_name}</td><td>${item.status}</td><td>${item.defect_count}</td><td><button data-id="${item.id}">标记关注</button></td></tr>`).join('');
  rows.querySelectorAll('button').forEach(button => button.onclick = () => update(button.dataset.id));
}
async function update(id) {
  const response = await fetch(`/api/v1/inspections/${id}/status`, {method:'POST', headers:{'Content-Type':'application/json'}, body:JSON.stringify({status:'attention'})});
  message.textContent = response.ok ? '状态已更新' : '更新失败';
  if (response.ok) load();
}
document.querySelector('#refresh').onclick = load;
load().catch(() => message.textContent = '无法读取巡检记录');
