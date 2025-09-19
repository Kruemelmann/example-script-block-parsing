# example-script-block-parsing


### Possible Use Cases for an `insertAdjacentHTML()` Flag to Control Script Execution

1. **Performance Optimization when Inserting HTML**
   - *Without execution:* Insert large HTML templates or UI components that contain scripts, but defer or skip their execution to improve rendering speed.
   - *With execution:* Dynamically load modules or widgets that should work immediately (e.g., a chart widget that brings its own script).

2. **Security / XSS Scenarios**
   - *Default (no execution):* Prevent scripts from running when inserting untrusted HTML from external sources or APIs.
   - *Explicit enable:* Allow scripts only for trusted HTML fragments, similar to a whitelisted module or widget.

3. **Lazy Execution / Progressive Enhancement**
   - Insert HTML containing scripts but keep them inactive initially.
   - Later activate them explicitly (re-insert or toggle execution) when user interaction requires it.

4. **Templating & Reuse**
   - Keep HTML snippets (with embedded scripts) as templates in the DOM without execution.
   - Reuse them by inserting with `executeScripts: true` in contexts where the scripts should actually run.

5. **Testing & Debugging**
   - Insert HTML with scripts but prevent execution to test DOM structure.
   - Useful for verifying markup correctness without triggering event handlers or side effects.

## Running the Service

### Installation
```bash
go mod download
```

### Running
```bash
go run main.go
```

The service will start on port 9090. You can access:
- **Web interface:** http://localhost:9090
- **API endpoints:**
  - `GET /api/scripts` - Get all script blocks
  - `GET /api/scripts/{id}` - Get specific script block by ID

