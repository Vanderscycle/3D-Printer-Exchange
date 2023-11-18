// Mock function to simulate fetching user data
const dummyUsers = {
    'john': { username: 'john', email: 'john@example.com', bio: 'Enthusiast in 3D printing.' },
    'jane': { username: 'jane', email: 'jane@example.com', bio: 'Lover of web development and design.' },
    'alex': { username: 'alex', email: 'alex@example.com', bio: 'Passionate about open source.' }
};

export async function fetchUserByUsername(username:any) {

    // Simulate an API call with a delay
    await new Promise(resolve => setTimeout(resolve, 1000));

    console.log('userData',dummyUsers[username]);
    return dummyUsers[username] || null;
}
