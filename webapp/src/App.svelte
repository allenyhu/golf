<script>
  let distance = '';
  let distances = [];

  function handleSubmit() {
    const dist = distance.trim();
    if (dist !== '' && !isNaN(dist) && parseFloat(dist) > 0) {
      distances = [...distances, dist];
      distance = '';
    }
  }

  function handleKeydown(event) {
    if (event.key === 'Enter') {
      event.preventDefault();
      handleSubmit();
    }
  }
</script>

<main>
  <div class="container">
    <h1>Distance from Hole</h1>
    
    <form on:submit|preventDefault={handleSubmit}>
      <div class="form-group">
        <label for="distance">Distance (yards)</label>
        <input
          id="distance"
          type="text"
          inputmode="numeric"
          bind:value={distance}
          placeholder="Enter distance"
          on:keydown={handleKeydown}
        />
      </div>
      <button type="submit">Save</button>
    </form>

    {#if distances.length > 0}
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>#</th>
              <th>Distance (yards)</th>
            </tr>
          </thead>
          <tbody>
            {#each distances as dist, index}
              <tr>
                <td>{index + 1}</td>
                <td>{dist}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
  </div>
</main>

<style>
  main {
    display: flex;
    justify-content: center;
    align-items: flex-start;
    min-height: 100vh;
    padding: 2rem;
    margin: 0;
  }

  .container {
    width: 100%;
    max-width: 600px;
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  h1 {
    font-size: 2.5rem;
    font-weight: bold;
    margin: 0;
    text-align: center;
    color: #333;
  }

  form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    background-color: #f8f9fa;
    padding: 1.5rem;
    border-radius: 8px;
  }

  .form-group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  label {
    font-weight: 600;
    color: #333;
  }

  input {
    padding: 0.75rem;
    font-size: 1rem;
    border: 1px solid #ddd;
    border-radius: 4px;
    outline: none;
  }

  input:focus {
    border-color: #007bff;
    box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
  }

  button {
    padding: 0.75rem 2rem;
    font-size: 1rem;
    font-weight: bold;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s;
  }

  button:hover {
    background-color: #0056b3;
  }

  .table-container {
    margin-top: 1rem;
    width: 100%;
  }

  table {
    width: 100%;
    border-collapse: collapse;
    background-color: white;
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    display: table;
  }

  thead {
    background-color: #007bff;
    color: white;
  }

  th {
    padding: 1rem;
    text-align: left;
    font-weight: 600;
  }

  tbody tr {
    border-bottom: 1px solid #eee;
  }

  tbody tr:last-child {
    border-bottom: none;
  }

  td {
    padding: 1rem;
    color: #333;
  }

  tbody tr:hover {
    background-color: #f8f9fa;
  }
</style>

