<script>
  let distance = '';
  let shotType = 'normal';
  
  // Load from localStorage synchronously before reactive statements run
  function loadFromLocalStorage() {
    if (typeof window === 'undefined') {
      return { holes: [], currentHole: 0 };
    }
    
    try {
      const savedHoles = localStorage.getItem('golfHoles');
      const savedCurrentHole = localStorage.getItem('golfCurrentHole');
      
      const holes = savedHoles ? JSON.parse(savedHoles) : [];
      const currentHole = savedCurrentHole ? parseInt(savedCurrentHole, 10) : 0;
      
      return { holes, currentHole };
    } catch (e) {
      console.error('Error loading from localStorage:', e);
      return { holes: [], currentHole: 0 };
    }
  }
  
  const { holes: initialHoles, currentHole: initialCurrentHole } = loadFromLocalStorage();
  let holes = initialHoles;
  let currentHole = initialCurrentHole;

  // Save to localStorage whenever holes or currentHole changes
  $: {
    if (typeof window !== 'undefined') {
      localStorage.setItem('golfHoles', JSON.stringify(holes));
      localStorage.setItem('golfCurrentHole', currentHole.toString());
    }
  }

  function handleSubmit() {
    const dist = distance.trim();
    if (dist !== '' && !isNaN(dist) && parseFloat(dist) > 0) {
      // Ensure we have a hole for the current index
      if (!holes[currentHole]) {
        holes = [...holes, []];
      }
      // Add shot to current hole, append 'b' for bunker or 'r' for recovery
      let shotValue = dist;
      if (shotType === 'bunker') {
        shotValue = dist + 'b';
      } else if (shotType === 'recovery') {
        shotValue = dist + 'r';
      } else if (shotType === 'putt') {
        shotValue = dist + 'p';
      }
      holes[currentHole] = [...holes[currentHole], shotValue];
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

  function addPenaltyStroke() {
    // Ensure we have a hole for the current index
    if (!holes[currentHole]) {
      holes = [...holes, []];
    }
    // Add 'pen' to current hole
    holes[currentHole] = [...holes[currentHole], 'pen'];
    holes = [...holes]; // Trigger reactivity
  }

  function formatShots(hole) {
    return hole.length > 0 ? hole.join(', ') : '-';
  }

  async function exportRound() {
    if (holes.length === 0) {
      return;
    }
    
    const lines = holes.map((hole, index) => {
      const holeNumber = index + 1;
      const shots = hole.length > 0 ? hole.join(', ') : '';
      return `hole ${holeNumber}: ${shots}`;
    });
    
    const exportText = lines.join('\n');
    
    try {
      await navigator.clipboard.writeText(exportText);
      alert('Round data copied to clipboard!');
    } catch (err) {
      console.error('Failed to copy to clipboard:', err);
      alert('Failed to copy to clipboard. Please try again.');
    }
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
      <div class="form-group">
        <label>Shot Type</label>
        <div class="toggle-group">
          <label class="toggle-option">
            <input type="radio" bind:group={shotType} value="normal" />
            <span>Normal</span>
          </label>
          <label class="toggle-option">
            <input type="radio" bind:group={shotType} value="putt" />
            <span>Putt</span>
          </label>
          <label class="toggle-option">
            <input type="radio" bind:group={shotType} value="recovery" />
            <span>Recovery</span>
          </label>
          <label class="toggle-option">
            <input type="radio" bind:group={shotType} value="bunker" />
            <span>Bunker</span>
          </label>
        </div>
      </div>
      <div class="button-group">
        <button type="button" class="penalty-button" on:click={addPenaltyStroke}>Penalty Stroke</button>
        <button type="submit">Save Shot</button>
      </div>
      <button type="button" on:click={nextHole}>Next Hole</button>
    </form>

    {#if holes.length > 0}
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>Hole</th>
              <th>Shots (yards)</th>
              <th>Score</th>
            </tr>
          </thead>
          <tbody>
            {#each holes as hole, holeIndex}
              <tr>
                <td class="hole-number">{holeIndex + 1}</td>
                <td>{formatShots(hole)}</td>
                <td class="score">{hole.length}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
      <button type="button" class="export-button" on:click={exportRound}>Export Round</button>
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

  .toggle-group {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
  }

  .toggle-option {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    cursor: pointer;
    padding: 0.5rem 1rem;
    border: 2px solid #ddd;
    border-radius: 4px;
    background-color: white;
    transition: all 0.3s;
    flex: 1;
    min-width: calc(50% - 0.25rem);
  }

  .toggle-option:hover {
    border-color: #007bff;
    background-color: #f0f8ff;
  }

  .toggle-option input[type="radio"] {
    margin: 0;
    cursor: pointer;
  }

  .toggle-option input[type="radio"]:checked + span {
    font-weight: bold;
    color: #007bff;
  }

  .toggle-option:has(input[type="radio"]:checked) {
    border-color: #007bff;
    background-color: #e7f3ff;
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

  button[type="button"]:not(.penalty-button) {
    background-color: #28a745;
  }

  button[type="button"]:not(.penalty-button):hover {
    background-color: #218838;
  }

  .penalty-button {
    background-color: #dc3545 !important;
  }

  .penalty-button:hover {
    background-color: #c82333 !important;
  }

  button[type="button"]:not(.penalty-button):not(.export-button) {
    width: 100%;
    margin-top: 0.5rem;
  }

  .export-button {
    width: 100%;
    background-color: #6c757d;
    margin-top: 1rem;
  }

  .export-button:hover {
    background-color: #5a6268;
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
    text-align: center;
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
    text-align: center;
  }

  .hole-number {
    font-weight: bold;
    background-color: #f8f9fa;
    color: #007bff;
  }

  .score {
    font-weight: bold;
    text-align: center;
    color: #28a745;
  }

  tbody tr:hover {
    background-color: #f8f9fa;
  }

  @media (max-width: 768px) {
    main {
      padding: 1rem;
    }

    .container {
      gap: 1.5rem;
    }

    .toggle-option {
      min-width: calc(50% - 0.25rem);
      padding: 0.5rem 0.75rem;
      font-size: 0.9rem;
    }

    form {
      padding: 1rem;
    }
  }
</style>

