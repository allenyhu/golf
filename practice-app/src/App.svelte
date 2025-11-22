<script>
  import { shapes, clubs, generateRandomNumber as generateShot } from './utils.js';

  let randomYardage = null;
  let randomShape = null;
  let randomClub = null;
  let mode = 'yardage';

  function generateRandomNumber() {
    const result = generateShot(mode, clubs, shapes);
    randomYardage = result.randomYardage;
    randomShape = result.randomShape;
    randomClub = result.randomClub;
  }
  
  function setMode(newMode) {
    mode = newMode;
    randomYardage = null;
    randomShape = null;
    randomClub = null;
  }
</script>

<main>
  <h1>Practice App</h1>
  
  <div class="mode-selector">
    <button 
      class="mode-button"
      class:active={mode === 'yardage'}
      on:click={() => setMode('yardage')}
    >
      Yardage
    </button>
    <button 
      class="mode-button"
      class:active={mode === 'yardage-shape'}
      on:click={() => setMode('yardage-shape')}
    >
      Yardage + Shape
    </button>
    <button 
      class="mode-button"
      class:active={mode === 'club-shape'}
      on:click={() => setMode('club-shape')}
    >
      Club + Shape
    </button>
    <button 
      class="mode-button"
      class:active={mode === 'club-yardage-shape'}
      on:click={() => setMode('club-yardage-shape')}
    >
      Club + Yardage + Shape
    </button>
  </div>

  {#if mode === 'yardage'}
    {#if randomYardage !== null}
      <p class="random-number">Yardage: <strong>{randomYardage}</strong></p>
    {:else}
      <p>Click the button to generate a random yardage</p>
    {/if}
  {:else if mode === 'yardage-shape'}
    {#if randomYardage !== null && randomShape !== null}
      <div class="random-display">
        <p class="random-number">Yardage: <strong>{randomYardage}</strong></p>
        <p class="random-shape">Shape: <strong>{randomShape}</strong></p>
      </div>
    {:else}
      <p>Click the button to generate a random yardage and shape</p>
    {/if}
  {:else if mode === 'club-shape'}
    {#if randomShape !== null && randomClub !== null}
      <div class="random-display">
        <p class="random-club">Club: <strong>{randomClub}</strong></p>
        <p class="random-shape">Shape: <strong>{randomShape}</strong></p>
      </div>
    {:else}
      <p>Click the button to generate a random club and shape</p>
    {/if}
  {:else if mode === 'club-yardage-shape'}
    {#if randomYardage !== null && randomShape !== null && randomClub !== null}
      <div class="random-display">
        <p class="random-club">Club: <strong>{randomClub}</strong></p>
        <p class="random-number">Yardage: <strong>{randomYardage}</strong></p>
        <p class="random-shape">Shape: <strong>{randomShape}</strong></p>
      </div>
    {:else}
      <p>Click the button to generate a random yardage, shape, and club</p>
    {/if}
  {/if}
  
  <button on:click={generateRandomNumber}>Generate Shot</button>
</main>

