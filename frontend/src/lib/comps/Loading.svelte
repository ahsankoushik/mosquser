<script lang="ts">
    import { ShieldPlus } from "@lucide/svelte";
    const {text=""}:{text:string} = $props();
</script>


<!-- Overlay -->
<div class="fixed inset-0 z-50 flex items-center justify-center bg-black/50">
    <!-- Loading Container -->
    <div class="flex flex-col items-center gap-4 bg-white p-6 rounded-lg shadow-lg">
      
      <!-- Optional Text -->
      <div class="flex items-center gap-2 text-xl font-semibold ">
        <ShieldPlus size={48} color="black" />
        <span>MosqUser</span>
      </div>
      <!-- Spinner -->
      <div class="animate-spin rounded-full h-12 w-12 border-4 border-black border-t-transparent"></div>
      {#if text != ""}
        {text}
      {/if}
    </div>
  </div>
  
