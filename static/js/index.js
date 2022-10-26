const todos = document.querySelectorAll('[data-todo]');

Array.from(todos).map((todo) => {
    todo.addEventListener('click', () => {
        const id = todo.dataset.todo;
        console.log(id);
        
        try {
            fetch("/todo/status", {
                headers: {
                    'Content-Type': 'application/json'
                },
                method: "POST",
                body: JSON.stringify({ id: id })
            })
            .then(res => {
                if (res.ok) {
                    todo.classList.toggle('done');
                }
            });   

        } catch (err) {
            console.error(err);
        }
    })
});