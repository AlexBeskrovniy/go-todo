const todos = document.querySelectorAll('[data-todo]');

Array.from(todos).map((todo) => {
    todo.addEventListener('click', () => {
        console.log(todo);
        const id = todo.dataset.todo;
        console.log(id);
        const body = {
            id: id
        }

        try {
            fetch("/todo/status", {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                method: "PUT",
                body: body,
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