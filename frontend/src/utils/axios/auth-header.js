export default function authHeader() {
  const token_id = localStorage.getItem("token");

  if (token_id) {
    return { Authorization: "Bearer " + token_id };
  }
  return {};
}
