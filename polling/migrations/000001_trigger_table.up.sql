-- Function to renumber questions after deletion
CREATE OR REPLACE FUNCTION renumber_questions() RETURNS TRIGGER AS $$
BEGIN
    WITH ordered_questions AS (
        SELECT id, ROW_NUMBER() OVER (PARTITION BY poll_id ORDER BY num) AS new_num
        FROM questions
    )
    UPDATE questions
    SET num = ordered_questions.new_num
    FROM ordered_questions
    WHERE questions.id = ordered_questions.id;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Function to renumber polls after deletion
CREATE OR REPLACE FUNCTION renumber_polls() RETURNS TRIGGER AS $$
BEGIN
    WITH ordered_polls AS (
        SELECT id, ROW_NUMBER() OVER (ORDER BY poll_num) AS new_num
        FROM polls
    )
    UPDATE polls
    SET poll_num = ordered_polls.new_num
    FROM ordered_polls
    WHERE polls.id = ordered_polls.id;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Conditional creation of the after_question_delete trigger
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 
        FROM pg_trigger 
        WHERE tgname = 'after_question_delete'
    ) THEN
        CREATE TRIGGER after_question_delete
        AFTER DELETE ON questions
        FOR EACH STATEMENT
        EXECUTE FUNCTION renumber_questions();
    END IF;
END
$$;

-- Conditional creation of the after_poll_delete trigger
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 
        FROM pg_trigger 
        WHERE tgname = 'after_poll_delete'
    ) THEN
        CREATE TRIGGER after_poll_delete
        AFTER DELETE ON polls
        FOR EACH STATEMENT
        EXECUTE FUNCTION renumber_polls();
    END IF;
END
$$;