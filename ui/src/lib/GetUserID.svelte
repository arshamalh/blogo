<script>
    // GET("api/v1/users/id", middlewares.RequireLogin, controllers.UserID)
    import { BaseURL } from "../stores";

    let msg = ""
    let status = 0;
    function GetUser() {
        fetch($BaseURL + "/users/id", {
          method: "GET",
          headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
          },
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


<section class="get_user_id">
    <button on:click="{GetUser}">Get User</button>
    {#if msg !== ""}
        <div>{msg}</div>
    {/if}
    {#if status !== 0}
        <div>{status}</div>
    {/if}
</section>

<style>
    .get_user_id {
        width: 90%;
        margin: 20px auto;
        height: 100px;
    }
</style>