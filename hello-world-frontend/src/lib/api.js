export async function getGreeting() {
  const res = await fetch("/api/greeting"); // no localhost or port!
  if (!res.ok) throw new Error("Failed to fetch greeting");
  const data = await res.json();
  return data.greeting;
}
