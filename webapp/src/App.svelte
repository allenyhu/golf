<script>
  let distance = '';
  let holes = [];
  let currentHole = 0;

  function handleSubmit() {
    const dist = distance.trim();
    if (dist !== '' && !isNaN(dist) && parseFloat(dist) > 0) {
      // Ensure we have a hole for the current index
      if (!holes[currentHole]) {
        holes = [...holes, []];
      }
      // Add shot to current hole
      holes[currentHole] = [...holes[currentHole], dist];
      holes = [...holes]; // Trigger reactivity
      distance = '';
    }
  }

  function nextHole() {
    // Ensure current hole exists before moving to next
    if (!holes[currentHole]) {
      holes = [...holes, []];
    }
    currentHole = holes.length;
    holes = [...holes, []];
    distance = '';
  }

  function handleKeydown(event) {
    if (event.key === 'Enter') {
      event.preventDefault();
      handleSubmit();
    }
  }

  function formatShots(hole) {
    return hole.length > 0 ? hole.join(', ') : '-';
  }
</script>

<main>
  <div class="container">
    <h1>Golf Round Tracker</h1>
    
    <div class="current-hole">
      <p>Current Hole: <strong>{currentHole + 1}</strong></p>
    </div>
    
    <form on:submit|preventDefault={handleSubmit}>
      <div class="form-group">
        <label for="distance">Shot Distance (yards)</label>
        <input
          id="distance"
          type="text"
          inputmode="numeric"
          bind:value={distance}
          placeholder="Enter distance"
          on:keydown={handleKeydown}
        />
      </div>
      <div class="button-group">
        <button type="submit">Save</button>
        <button type="button" on:click={nextHole}>Next Hole</button>
      </div>
    </form>

    {#if holes.length > 0}
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>Hole</th>
              <th>Shots (yards)</th>
            </tr>
          </thead>
          <tbody>
            {#each holes as hole, holeIndex}
              <tr>
                <td class="hole-number">{holeIndex + 1}</td>
                <td>{formatShots(hole)}</td>
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

  .current-hole {
    text-align: center;
    font-size: 1.2rem;
    color: #333;
  }

  .current-hole strong {
    color: #007bff;
    font-size: 1.5rem;
  }

  .button-group {
    display: flex;
    gap: 1rem;
  }

  button {
    flex: 1;
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

  button[type="button"] {
    background-color: #28a745;
  }

  button[type="button"]:hover {
    background-color: #218838;
  }

  .table-container {
    margin-top: 1rem;
    width: 100%;
    overflow-x: auto;
  }

  table {
    width: 100%;
    min-width: 300px;
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

  .hole-number {
    font-weight: bold;
    background-color: #f8f9fa;
    color: #007bff;
  }

  tbody tr:hover {
    background-color: #f8f9fa;
  }
</style>

