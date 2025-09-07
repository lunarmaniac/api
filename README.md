# api.dorocha.dev

minimal api, right now simply serving coding tips to progress my knowledge of golang.

## endpoints

- `/tip` — returns a single random coding tip as JSON.  
  each tip also includes its category for context.  

  example response:
  ```json
  {
    "category": "Best Practices",
    "tip": "Use meaningful variable names."
  }
    ````

## categories

tips are organized into the following categories internally:

* `Best Practices`
* `Debugging & Testing`
* `Performance & Optimization`
* `Tools & Workflow`

currently, `/tip` returns a tip randomly from **all categories** combined.