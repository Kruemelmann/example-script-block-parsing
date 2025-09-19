document.addEventListener('DOMContentLoaded', function() {

    async function loadScript(scriptId) {
        try {
            const response = await fetch(`/api/scripts/${scriptId}`);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const scriptData = await response.json();
            if (scriptData.error) {
                console.error('Script error:', scriptData.error);
                return;
            }

            const scriptElement = document.createElement('script');
            scriptElement.type = 'text/javascript';
            scriptElement.innerHTML = scriptData.content;

            document.head.insertAdjacentHTML('beforeend', scriptElement.outerHTML);

            console.log(`Script ${scriptData.id} loaded`);

        } catch (error) {
            console.error('Error loading script:', error);
        }
    }

    window.loadScript = loadScript;
});
