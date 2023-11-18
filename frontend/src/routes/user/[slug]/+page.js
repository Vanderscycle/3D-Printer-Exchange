import { fetchUserByUsername } from '$lib/api/userData';

export async function load({ params }) {
  const username  = params.slug;
  console.log('username', username);
  const user = await fetchUserByUsername(username);
  return { props: { user } };
}
