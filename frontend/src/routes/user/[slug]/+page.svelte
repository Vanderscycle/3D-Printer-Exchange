<script>
export let data;
console.log(data);
let user = data.props.user
if (!user) {
    user = { username: 'Not found', email: 'N/A', bio: 'User does not exist.' };
}
</script>

<div class="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4">
    <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <h1 class="text-2xl font-bold mb-2">User Profile</h1>
        <p><strong>Username:</strong> {user.username}</p>
        <p><strong>Email:</strong> {user.email}</p>
        <p><strong>Bio:</strong> {user.bio}</p>
    </div>
</div>

<style>
/* Additional styling if needed */
</style>
