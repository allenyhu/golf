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
  let editingHoleIndex = null;
  let showEditPopup = false;
  let hoveredHoleIndex = null;
  let editedShots = [];
  let newShotValue = '';

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

  function openEditPopup(holeIndex) {
    editingHoleIndex = holeIndex;
    // Create a copy of the shots array for editing
    editedShots = holes[holeIndex] ? [...holes[holeIndex]] : [];
    showEditPopup = true;
  }

  function closeEditPopup() {
    showEditPopup = false;
    editingHoleIndex = null;
    editedShots = [];
    newShotValue = '';
  }

  function saveEditedShots() {
    if (editingHoleIndex !== null && holes[editingHoleIndex]) {
      // Save the edited shots back to the holes array
      holes[editingHoleIndex] = [...editedShots];
      holes = [...holes]; // Trigger reactivity
      closeEditPopup();
    }
  }

  function addNewShot() {
    const trimmed = newShotValue.trim();
    if (trimmed !== '') {
      editedShots = [...editedShots, trimmed];
      editedShots = [...editedShots]; // Trigger reactivity
      newShotValue = '';
    }
  }

  function deleteShot(shotIndex) {
    if (shotIndex >= 0 && shotIndex < editedShots.length) {
      editedShots = editedShots.filter((_, i) => i !== shotIndex);
      editedShots = [...editedShots]; // Trigger reactivity
    }
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
                <td 
                  class="shots-cell"
                  on:click={() => openEditPopup(holeIndex)}
                  on:mouseenter={() => hoveredHoleIndex = holeIndex}
                  on:mouseleave={() => hoveredHoleIndex = null}
                >
                  <span>{formatShots(hole)}</span>
                  {#if hoveredHoleIndex === holeIndex}
                    <button 
                      class="edit-button-small"
                      on:click|stopPropagation={() => openEditPopup(holeIndex)}
                    >
                      Edit
                    </button>
                  {/if}
                </td>
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

{#if showEditPopup && editingHoleIndex !== null}
  <div class="popup-overlay" on:click={closeEditPopup}>
    <div class="popup-content" on:click|stopPropagation>
      <div class="popup-header">
        <h2>Edit Shots - Hole {editingHoleIndex + 1}</h2>
        <button class="close-button" on:click={closeEditPopup}>×</button>
      </div>
      <div class="popup-body">
        {#if editedShots.length === 0}
          <p>No shots recorded for this hole.</p>
        {:else}
          <div class="shots-list">
            {#each editedShots as shot, shotIndex}
              <div class="shot-item">
                <input
                  type="text"
                  class="shot-input"
                  bind:value={editedShots[shotIndex]}
                  placeholder="Shot value"
                />
                <button 
                  class="delete-shot-button"
                  on:click={() => deleteShot(shotIndex)}
                >
                  Delete
                </button>
              </div>
            {/each}
          </div>
        {/if}
        <div class="add-shot-section">
          <div class="add-shot-input-group">
            <input
              type="text"
              class="new-shot-input"
              bind:value={newShotValue}
              placeholder="Enter new shot value"
              on:keydown={(e) => e.key === 'Enter' && addNewShot()}
            />
            <button 
              class="add-shot-button"
              on:click={addNewShot}
            >
              Add Shot
            </button>
          </div>
        </div>
      </div>
      <div class="popup-footer">
        <button class="save-button" on:click={saveEditedShots}>Save</button>
        <button class="cancel-button" on:click={closeEditPopup}>Cancel</button>
      </div>
    </div>
  </div>
{/if}

