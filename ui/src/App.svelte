<script>
  import GetUserId from "./lib/GetUserID.svelte";
  import Login from "./lib/Login.svelte";
  import Logout from "./lib/Logout.svelte";
</script>

<main>
  <div class="box">
    <GetUserId />
    <Login />
    <Logout />
  </div>
</main>

<style>
  :root {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
      Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  }

  main {
    text-align: center;
    padding: 1em;
    margin: 0 auto;
    display: flex;
    justify-content: center;
  }

  .box {
    background-color: #e0e0ee;
    width: 500px;
    height: 550px;
  }
</style>
