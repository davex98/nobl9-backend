const apiUrl = `${Cypress.env("apiUrl")}`

describe('Nobl9-backend', () => {
    it('should get an error, as the request param is invalid', () => {
        cy.request({
            failOnStatusCode: false,
            method: 'GET',
            url: `${apiUrl}/random/mean?requests=0&length=2`,
        }).then((response) => {
            expect(response.status).to.eq(400)
        })
    })

    it('should get an error, as the length param is invalid', () => {
        cy.request({
            failOnStatusCode: false,
            method: 'GET',
            url: `${apiUrl}/random/mean?requests=1&length=0`,
        }).then((response) => {
            expect(response.status).to.eq(400)
        })
    })

    it('should get a valid response for one number', () => {
        cy.request({
            failOnStatusCode: false,
            method: 'GET',
            url: `${apiUrl}/random/mean?requests=1&length=1`,
        }).then((response) => {
            expect(response.status).to.eq(200)
            expect(response.body.length).to.eq(2)
            expect(response.body[0].stddev).to.eq(0)
            expect(response.body[1].stddev).to.eq(0)
        })
    })

    it('should get a valid response for 2 requests', () => {
        cy.request({
            failOnStatusCode: false,
            method: 'GET',
            url: `${apiUrl}/random/mean?requests=2&length=1`,
        }).then((response) => {
            expect(response.status).to.eq(200)
            expect(response.body.length).to.eq(3)
            expect(response.body[0].stddev).to.eq(0)
            expect(response.body[1].stddev).to.eq(0)
        })
    })

    it('should get a valid response for 2 requests with different length', () => {
        cy.request({
            failOnStatusCode: false,
            method: 'GET',
            url: `${apiUrl}/random/mean?requests=2&length=3`,
        }).then((response) => {
            expect(response.status).to.eq(200)
            expect(response.body.length).to.eq(3)
        })
    })
})