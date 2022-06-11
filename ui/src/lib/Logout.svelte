<script>
    // POST("api/v1/users/logout", middlewares.RequireLogin, controllers.UserLogout)
    import { BaseURL } from "../stores";

    let msg = "";
    let status = 0;

function Logout() {
    fetch($BaseURL + "/users/logout", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      credentials: "include",
    })
    .then(res => {
        status = res.status;
        if (res.status === 200) {
            msg = "logout Successful";
        } else if (res.status === 401) {
            msg = "you are already logged out";
        }
    })
}
</script>

<section class="logout">
    <button on:click="{Logout}">Logout</button>
    {#if msg !== ""}
        <div>{msg}</div>
    {/if}
    {#if status !== 0}
        <div>{status}</div>
    {/if}
</section>

<style>
    .logout {
        width: 90%;
        margin: 20px auto;
        height: 100px;
    }
</style>