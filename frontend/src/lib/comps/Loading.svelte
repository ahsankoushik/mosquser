<script>
    import { ShieldX } from "@lucide/svelte";

</script>


<!-- Overlay -->
<div class="fixed inset-0 z-50 flex items-center justify-center bg-black/50">
    <!-- Loading Container -->
    <div class="flex flex-col items-center gap-4 bg-white p-6 rounded-lg shadow-lg">
      
      <!-- Optional Text -->
      <div class="flex items-center gap-2 text-xl font-semibold ">
        <ShieldX size={48} color="purple" />
        <span>MosqUser</span>
      </div>
      <!-- Spinner -->
      <div class="animate-spin rounded-full h-12 w-12 border-4 border-purple-500 border-t-transparent"></div>
    </div>
  </div>
  