<script>
    // POST("api/v1/users/login", controllers.UserLogin)
    import { BaseURL } from "../stores";

    let username = "";
    let password = "";
    let status = 0;
    let msg = "";

    function Login() {
        fetch($BaseURL + "/users/login", {
            method: "POST",
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ "username": username, "password": password }),
            credentials: "include",
        })
        .then(res => {
            status = res.status;
            return res.json()
        })
        .then(user => {
            msg = JSON.stringify(user)
        });
    }
</script>


<section class="login">
    <input type="text" name="username" placeholder="username" bind:value={username}>
    <input type="text" name="password" placeholder="password" bind:value={password}>
    <button on:click="{Login}">Login</button>
    {#if msg !== ""}
        <div>{msg}</div>
    {/if}
    {#if status !== 0}
        <div>{status}</div>
    {/if}
</section>

<section class="get_user_id">
</section>

<style>
    .get_user_id {
        width: 90%;
        margin: 20px auto;
        height: 100px;
    }
</style>